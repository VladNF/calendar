package serverhttp

import (
	"net/http"
)

func (s *HTTPServer) NoEror(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (s *HTTPServer) InternalError(err error, w http.ResponseWriter, r *http.Request) {
	s.httpRespondWithError(err, w, r, "Internal server error", http.StatusInternalServerError)
}

func (s *HTTPServer) BadRequest(err error, w http.ResponseWriter, r *http.Request) {
	s.httpRespondWithError(err, w, r, "Bad request", http.StatusBadRequest)
}

func (s *HTTPServer) NotFound(err error, w http.ResponseWriter, r *http.Request) {
	s.httpRespondWithError(err, w, r, "Not found", http.StatusNotFound)
}

func (s *HTTPServer) httpRespondWithError(err error, w http.ResponseWriter, _ *http.Request, logMsg string, status int) {
	s.log.Errorf("%v - %v", logMsg, err)
	w.WriteHeader(status)
}
