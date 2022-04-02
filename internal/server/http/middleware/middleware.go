package middleware

import (
	"net"
	"net/http"
	"time"

	"github.com/VladNF/otus-golang/hw12_13_14_15_calendar/internal/common"
)

type loggingRW struct {
	http.ResponseWriter
	status int
}

func (w *loggingRW) WriteHeader(statusCode int) {
	w.status = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func LoggingMiddleware(log common.Logger) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			begin := time.Now()
			logWriter := &loggingRW{w, http.StatusOK}
			h.ServeHTTP(logWriter, r)
			ip, _, err := net.SplitHostPort(r.RemoteAddr)
			if err != nil {
				ip = r.RemoteAddr
			}
			log.Infof("%v [%v] %v %v %v %v %v %v", ip, begin.Format("02/Jan/2006:15:04:05 -0700"),
				r.Method, r.URL.EscapedPath(), r.Proto, logWriter.status, time.Since(begin), r.UserAgent())
		})
	}
}
