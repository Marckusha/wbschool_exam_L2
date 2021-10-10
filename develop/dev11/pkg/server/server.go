package server

import (
	"net/http"
)

type createEventHandler struct {
	serv      Service
	transport CreateEventTransport
}

func (s *createEventHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	event, userID, err := s.transport.DecodeRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = s.serv.CreateEvent(userID, event)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
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

func (s *updateEventHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	userID, evs, err := s.transport.DecodeRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(evs) != 2 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = s.serv.UpdateEvent(userID, evs[0], evs[1])
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}
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

	event, userID, err := s.transport.DecodeRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = s.serv.DeleteEvent(userID, event)

	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func NewDeleteHandler(s Service, t DeleteEventTransport) http.Handler {
	return &deleteEventHandler{
		serv:      s,
		transport: t,
	}
}

type getEventsForDay struct {
	serv      Service
	transport GetEventsDayTransport
}

func (s *getEventsForDay) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	event, userID, err := s.transport.DecodeRequest(r)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	notes, err := s.serv.EventsForDay(userID, event)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	if err = s.transport.EncodeResponse(r, w, notes); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func NewGetDayHandler(s Service, t GetEventsDayTransport) http.Handler {
	return &getEventsForDay{
		serv:      s,
		transport: t,
	}
}

type getEventsForWeek struct {
	serv      Service
	transport GetEventsWeekTransport
}

func (s *getEventsForWeek) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	event, userID, err := s.transport.DecodeRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	notes, err := s.serv.EventsForWeek(userID, event)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	if err = s.transport.EncodeResponse(r, w, notes); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func NewGetWeekHandler(s Service, t GetEventsWeekTransport) http.Handler {
	return &getEventsForWeek{
		serv:      s,
		transport: t,
	}
}

type getEventsForMonth struct {
	serv      Service
	transport GetEventsMonthTransport
}

func (s *getEventsForMonth) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	event, userID, err := s.transport.DecodeRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	notes, err := s.serv.EventsForMonth(userID, event)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	if err = s.transport.EncodeResponse(r, w, notes); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func NewGetMonthHandler(s Service, t GetEventsMonthTransport) http.Handler {
	return &getEventsForMonth{
		serv:      s,
		transport: t,
	}
}
