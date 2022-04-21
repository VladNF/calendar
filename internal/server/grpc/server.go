package servergrpc

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/VladNF/calendar/internal/app"
	"github.com/VladNF/calendar/internal/common"
	"github.com/VladNF/calendar/internal/models"
	"github.com/VladNF/calendar/internal/server/grpc/gen"
	"github.com/golang/protobuf/ptypes/empty"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func init() {
	grpc_logrus.ReplaceGrpcLogger(logrus.NewEntry(logrus.StandardLogger()))
}

type GRPCServer struct {
	app    *app.App
	host   string
	port   string
	log    common.Logger
	server *grpc.Server
}

func (s *GRPCServer) GetEvent(ctx context.Context, id *gen.EventId) (*gen.Event, error) {
	e, err := s.app.GetEvent(ctx, id.GetId())
	if err != nil {
		return nil, err
	}
	return &gen.Event{
		AlertBefore: int64(e.AlertBefore.Seconds()),
		EndsAt:      timestamppb.New(e.EndsAt),
		Id:          e.ID,
		Notes:       e.Notes,
		OwnerId:     e.OwnerID,
		StartsAt:    timestamppb.New(e.StartsAt),
		Title:       e.Title,
	}, nil
}

func (s *GRPCServer) PutEvent(ctx context.Context, event *gen.Event) (*gen.Event, error) {
	m, err := models.NewEvent(event.Id, event.Title, event.StartsAt.AsTime(), event.EndsAt.AsTime(), event.OwnerId)
	if err != nil {
		return nil, err
	}
	m.AlertBefore = time.Duration(event.AlertBefore * 1_000_000_000)
	m.Notes = event.Notes
	if err = s.app.UpdateEvent(ctx, m); err != nil {
		return nil, err
	}
	return event, nil
}

func (s *GRPCServer) DeleteEvent(ctx context.Context, id *gen.EventId) (*empty.Empty, error) {
	e, err := s.app.GetEvent(ctx, id.GetId())
	if err != nil {
		return nil, err
	}

	if err := s.app.DeleteEvent(ctx, e); err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

func (s *GRPCServer) ListEvents(ctx context.Context, request *gen.ListEventsRequest) (*gen.ListEventsResponse, error) {
	var events []*models.Event
	var err error
	switch request.Agenda {
	case gen.ListEventsRequest_DAILY:
		events, err = s.app.GetDailyAgenda(ctx, request.StartFrom.AsTime())
	case gen.ListEventsRequest_WEEKLY:
		events, err = s.app.GetWeeklyAgenda(ctx, request.StartFrom.AsTime())
	case gen.ListEventsRequest_MONTHLY:
		events, err = s.app.GetMonthlyAgenda(ctx, request.StartFrom.AsTime())
	default:
		return nil, fmt.Errorf("http: invalid agenda type %s", request.Agenda)
	}

	if err != nil {
		return nil, err
	}

	result := make([]*gen.Event, 0, len(events))
	for _, e := range events {
		result = append(result, &gen.Event{
			AlertBefore: int64(e.AlertBefore.Seconds()),
			EndsAt:      timestamppb.New(e.EndsAt),
			Id:          e.ID,
			Notes:       e.Notes,
			OwnerId:     e.OwnerID,
			StartsAt:    timestamppb.New(e.StartsAt),
			Title:       e.Title,
		})
	}
	return &gen.ListEventsResponse{Events: result}, nil
}

func NewServer(host string, port string, logger common.Logger, app *app.App) *GRPCServer {
	return &GRPCServer{host: host, port: port, app: app, log: logger}
}

func (s *GRPCServer) Start() error {
	addr := net.JoinHostPort(s.host, s.port)
	logrusEntry := logrus.NewEntry(logrus.StandardLogger())

	s.server = grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_logrus.UnaryServerInterceptor(logrusEntry),
		),
		grpc_middleware.WithStreamServerChain(
			grpc_ctxtags.StreamServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_logrus.StreamServerInterceptor(logrusEntry),
		),
	)
	gen.RegisterCalendarServiceServer(s.server, s)

	listen, err := net.Listen("tcp", addr)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.WithField("grpcEndpoint", addr).Info("Starting: gRPC Listener")
	logrus.Fatal(s.server.Serve(listen))
	return nil
}

func (s *GRPCServer) Stop(ctx context.Context) error {
	s.server.GracefulStop()
	return nil
}
