package servergrpc

import (
	"context"
	"net"
	"testing"
	"time"

	"github.com/VladNF/calendar/internal/app"
	"github.com/VladNF/calendar/internal/common"
	"github.com/VladNF/calendar/internal/server/grpc/gen"
	"github.com/VladNF/calendar/internal/storage"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const bufSize = 1024 * 1024

var (
	grpcListener *bufconn.Listener
	grpcServer   *GRPCServer
)

func init() {
	log := common.NewLogger("debug", "")
	grpcListener = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	grpcServer = makeServer()
	gen.RegisterCalendarServiceServer(s, grpcServer)
	go func() {
		if err := s.Serve(grpcListener); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return grpcListener.Dial()
}

func TestGrpcServer(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	tc := gen.NewCalendarServiceClient(conn)

	startDate := time.Date(2021, 1, 1, 8, 0, 0, 0, time.UTC)
	startTime := startDate.Add(8 * time.Hour)
	t.Run("Get Event", func(t *testing.T) {
		event, err := grpcServer.app.CreateEvent(ctx, "", "today event", startTime, startTime.Add(time.Hour), "test")
		require.NoError(t, err)

		expected := gen.Event{
			AlertBefore: int64(event.AlertBefore.Seconds()),
			EndsAt:      timestamppb.New(event.EndsAt),
			Id:          event.ID,
			Notes:       event.Notes,
			OwnerId:     event.OwnerID,
			StartsAt:    timestamppb.New(event.StartsAt),
			Title:       event.Title,
		}
		r, err := tc.GetEvent(ctx, &gen.EventId{Id: event.ID})
		require.NoError(t, err)
		require.Equal(t, r.String(), expected.String())

		_, err = tc.GetEvent(ctx, &gen.EventId{Id: "42"})
		require.NotNil(t, err)
	})

	t.Run("Delete Event", func(t *testing.T) {
		event, err := grpcServer.app.CreateEvent(ctx, "", "today event", startTime, startTime.Add(time.Hour), "test")
		require.NoError(t, err)

		_, err = tc.DeleteEvent(ctx, &gen.EventId{Id: event.ID})
		require.NoError(t, err)

		_, err = tc.GetEvent(ctx, &gen.EventId{Id: event.ID})
		require.NotNil(t, err)
	})

	t.Run("Update Event", func(t *testing.T) {
		event, err := grpcServer.app.CreateEvent(ctx, "", "today event", startTime, startTime.Add(time.Hour), "test")
		require.NoError(t, err)

		expected := gen.Event{
			AlertBefore: int64(event.AlertBefore.Seconds()),
			EndsAt:      timestamppb.New(event.EndsAt),
			Id:          event.ID,
			Notes:       "42",
			OwnerId:     event.OwnerID,
			StartsAt:    timestamppb.New(event.StartsAt),
			Title:       event.Title,
		}
		_, err = tc.PutEvent(ctx, &expected)
		require.NoError(t, err)

		r, err := tc.GetEvent(ctx, &gen.EventId{Id: event.ID})
		require.NoError(t, err)
		require.Equal(t, r.Notes, expected.Notes)
	})

	t.Run("List Events", func(t *testing.T) {
		startTime := startTime.AddDate(0, 1, 0)
		event, err := grpcServer.app.CreateEvent(
			ctx, "",
			"today event",
			startTime,
			startTime.Add(time.Hour),
			"test",
		)
		require.NoError(t, err)

		expected := gen.Event{
			AlertBefore: int64(event.AlertBefore.Seconds()),
			EndsAt:      timestamppb.New(event.EndsAt),
			Id:          event.ID,
			OwnerId:     event.OwnerID,
			StartsAt:    timestamppb.New(event.StartsAt),
			Title:       event.Title,
		}

		request := &gen.ListEventsRequest{
			Agenda:    gen.ListEventsRequest_DAILY,
			StartFrom: timestamppb.New(startDate.AddDate(0, 1, 0)),
		}
		r, err := tc.ListEvents(ctx, request)
		require.NoError(t, err)
		require.Len(t, r.Events, 1)
		require.Equal(t, r.Events[0].String(), expected.String())
	})
}

func makeServer() *GRPCServer {
	log := common.NewLogger("debug", "")

	eventsRepo, err := storage.NewStorage("in-memory")
	if err != nil {
		log.Fatalf("storage was not created: %v", err)
	}
	calendar := app.New(log, eventsRepo)
	return NewServer("", "", log, calendar)
}
