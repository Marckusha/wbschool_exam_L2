package server

import (
	"fmt"
	"net/http"
)

type createEventHandler struct {
	serv      Service
	transport CreateEventTransport
}

func (s *createEventHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "create method")

	event, err := s.transport.DecodeRequest(r)
	if err != nil {
		return
	}

	err = s.serv.CreateEvent(event)
	if err != nil {
		return
	}

	fmt.Println(event)
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

	fmt.Fprintln(w, "update method")

	event, err := s.transport.DecodeRequest(r)
	if err != nil {
		return
	}

	//TODO!!!
	err = s.serv.UpdateEvent(event, event)
	if err != nil {
		return
	}

	fmt.Println(event)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func NewUpdateHandler(s Service, t UpdateEventTransport) http.Handler {
	return &updateEventHandler{
		serv:      s,
		transport: t,
	}
}

type deleteEventHandler struct {
	serv      Service
	transport DeleteEventTransport
}

func (s *deleteEventHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "delete method")

	event, err := s.transport.DecodeRequest(r)
	if err != nil {
		return
	}

	fmt.Println(1)

	err = s.serv.DeleteEvent(event)

	fmt.Println(2)
	if err != nil {
		return
	}

	fmt.Println(3)
	fmt.Println(event)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func NewDeleteHandler(s Service, t DeleteEventTransport) http.Handler {
	return &deleteEventHandler{
		serv:      s,
		transport: t,
	}
}
