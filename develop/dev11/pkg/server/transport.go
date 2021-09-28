package server

import (
	"encoding/json"
	"io"
	"net/http"
)

type CreateEventTransport interface {
	DecodeRequest(req *http.Request) (event Event, err error)
	EncodeResponse(req *http.Request, res *http.Response) (err error)
}

type createEventTransport struct{}

func (ev *createEventTransport) DecodeRequest(req *http.Request) (event Event, err error) {

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
	DecodeRequest(req *http.Request) (event Event, err error)
	EncodeResponse(req *http.Request, res *http.Response) (err error)
}

type updateEventTransport struct{}

func (ev *updateEventTransport) DecodeRequest(req *http.Request) (event Event, err error) {

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

func (ev *updateEventTransport) EncodeResponse(req *http.Request, res *http.Response) (err error) {
	return
}

func NewUpdateEventTransport() UpdateEventTransport {
	return &updateEventTransport{}
}
