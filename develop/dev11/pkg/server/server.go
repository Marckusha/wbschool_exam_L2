package server

import (
	"net/http"
)

type createEventHandler struct {
	serv      Service
	transport CreateEventTransport
}

func (s *createEventHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	event, err := s.transport.DecodeRequest(r)
	if err != nil {
		return
	}

	err = s.serv.CreateEvent(event)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func NewCreateHandler(s Service, t CreateEventTransport) http.Handler {
	return &createEventHandler{
		serv:      s,
		transport: t,
	}
}

type updateEventHandler struct {
	serv      Service
	transport UpdateEventTransport
}

//TODO
func (s *updateEventHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	event, err := s.transport.DecodeRequest(r)
	if err != nil {
		return
	}

	err = s.serv.UpdateEvent(event, event)
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func NewUpdateHandler(s Service, t UpdateEventTransport) http.Handler {
	return &updateEventHandler{
		serv:      s,
		transport: t,
	}
}
