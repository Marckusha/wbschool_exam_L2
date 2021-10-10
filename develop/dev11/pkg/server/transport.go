package server

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type CreateEventTransport interface {
	DecodeRequest(req *http.Request) (event Event, userID int, err error)
	EncodeResponse(req *http.Request, res *http.Response) (err error)
}

type resultOut struct {
	Result string `json:result`
}

type createEventTransport struct{}

func (ev *createEventTransport) DecodeRequest(req *http.Request) (event Event, userID int, err error) {

	userID, err = strconv.Atoi(string(req.Header.Get("user_id")))
	if err != nil {
		return
	}

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return
	}

	var request Event
	if err = json.Unmarshal(body, &request); err != nil {
		return
	}

	event = request

	return
}

func (ev *createEventTransport) EncodeResponse(req *http.Request, res *http.Response) (err error) {
	return
}

func NewCreateEventTransport() CreateEventTransport {
	return &createEventTransport{}
}

type UpdateEventTransport interface {
	DecodeRequest(req *http.Request) (userID int, event []Event, err error)
	EncodeResponse(req *http.Request, res *http.Response) (err error)
}

type updateEventTransport struct{}

func (ev *updateEventTransport) DecodeRequest(req *http.Request) (userID int, event []Event, err error) {

	userID, err = strconv.Atoi(string(req.Header.Get("user_id")))
	if err != nil {
		return
	}

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return
	}

	var request []Event
	if err = json.Unmarshal(body, &request); err != nil {
		return
	}

	event = request

	return
}

func (ev *updateEventTransport) EncodeResponse(req *http.Request, res *http.Response) (err error) {
	return
}

func NewUpdateEventTransport() UpdateEventTransport {
	return &updateEventTransport{}
}

type DeleteEventTransport interface {
	DecodeRequest(req *http.Request) (event Event, userID int, err error)
	EncodeResponse(req *http.Request, res *http.Response) (err error)
}

type deleteEventTransport struct{}

func (ev *deleteEventTransport) DecodeRequest(req *http.Request) (event Event, userID int, err error) {

	userID, err = strconv.Atoi(string(req.Header.Get("user_id")))
	if err != nil {
		return
	}

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return
	}

	var request Event
	if err = json.Unmarshal(body, &request); err != nil {
		return
	}

	event = request

	return
}

func (ev *deleteEventTransport) EncodeResponse(req *http.Request, res *http.Response) (err error) {
	return
}

func NewDeleteTransport() DeleteEventTransport {
	return &deleteEventTransport{}
}

type GetEventsDayTransport interface {
	DecodeRequest(req *http.Request) (event Event, userID int, err error)
	EncodeResponse(req *http.Request, res http.ResponseWriter, notes []string) (err error)
}

type getEventsDayTransport struct{}

func (ev *getEventsDayTransport) DecodeRequest(req *http.Request) (event Event, userID int, err error) {
	if req.Method == http.MethodGet {
		date := req.URL.Query().Get("date")
		userID, err = strconv.Atoi(req.URL.Query().Get("user_id"))
		if err != nil {
			return
		}

		event, err = NewEvent(date)
		if err != nil {
			return
		}
	}

	return
}

func (ev *getEventsDayTransport) EncodeResponse(req *http.Request, res http.ResponseWriter, notes []string) (err error) {
	res.Header().Set("Content-Type", "application/json")

	body, err := json.Marshal(resultOut{
		Result: strings.Join(notes, ", "),
	})

	if err != nil {
		res.Write([]byte(err.Error()))
		return
	}

	_, err = res.Write(body)
	if err != nil {
		res.Write([]byte(err.Error()))
		return
	}

	return
}

func NewGetEventsDayTransport() GetEventsDayTransport {
	return &getEventsDayTransport{}
}

type GetEventsWeekTransport interface {
	DecodeRequest(req *http.Request) (event Event, userID int, err error)
	EncodeResponse(req *http.Request, res http.ResponseWriter, notes []string) (err error)
}

type getEventsWeekTransport struct{}

func (ev *getEventsWeekTransport) DecodeRequest(req *http.Request) (event Event, userID int, err error) {

	if req.Method == http.MethodGet {
		date := req.URL.Query().Get("date")
		userID, err = strconv.Atoi(req.URL.Query().Get("user_id"))
		if err != nil {
			return
		}

		event, err = NewEvent(date)
		if err != nil {
			return
		}
	}

	return
}

func (ev *getEventsWeekTransport) EncodeResponse(req *http.Request, res http.ResponseWriter, notes []string) (err error) {
	res.Header().Set("Content-Type", "application/json")

	body, err := json.Marshal(resultOut{
		Result: strings.Join(notes, ", "),
	})
	if err != nil {
		res.Write([]byte(err.Error()))
		return
	}

	_, err = res.Write(body)
	if err != nil {
		res.Write([]byte(err.Error()))
		return
	}

	return
}

func NewGetEventsWeekTransport() GetEventsWeekTransport {
	return &getEventsWeekTransport{}
}

type GetEventsMonthTransport interface {
	DecodeRequest(req *http.Request) (event Event, userID int, err error)
	EncodeResponse(req *http.Request, res http.ResponseWriter, notes []string) (err error)
}

type getEventsMonthTransport struct{}

func (ev *getEventsMonthTransport) DecodeRequest(req *http.Request) (event Event, userID int, err error) {
	if req.Method == http.MethodGet {
		date := req.URL.Query().Get("date")
		userID, err = strconv.Atoi(req.URL.Query().Get("user_id"))
		if err != nil {
			return
		}

		event, err = NewEvent(date)
		if err != nil {
			return
		}
	}

	return
}

func (ev *getEventsMonthTransport) EncodeResponse(req *http.Request, res http.ResponseWriter, notes []string) (err error) {
	res.Header().Set("Content-Type", "application/json")

	body, err := json.Marshal(resultOut{
		Result: strings.Join(notes, ", "),
	})
	if err != nil {
		res.Write([]byte(err.Error()))
		return
	}

	_, err = res.Write(body)
	if err != nil {
		res.Write([]byte(err.Error()))
		return
	}

	return
}

func NewGetEventsMonthTransport() GetEventsMonthTransport {
	return &getEventsMonthTransport{}
}
