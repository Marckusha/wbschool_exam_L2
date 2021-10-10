package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"wbschool_exam_L2/develop/dev11/pkg/server"

	"github.com/sirupsen/logrus"
)

func main() {
	body, err := ioutil.ReadFile("config.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
		return
	}

	port := string(body)

	logger := logrus.New()
	serv := server.NewService()
	serv = server.NewLogMiddleware(serv, logger)

	createTransport := server.NewCreateEventTransport()
	updateTransport := server.NewUpdateEventTransport()
	deleteTransport := server.NewDeleteTransport()
	getEventsDayTransport := server.NewGetEventsDayTransport()
	getEventsWeekTransport := server.NewGetEventsWeekTransport()
	getEventsMonthTransport := server.NewGetEventsMonthTransport()

	http.Handle("/create_event", server.NewCreateHandler(serv, createTransport))
	http.Handle("/update_event", server.NewUpdateHandler(serv, updateTransport))
	http.Handle("/delete_event", server.NewDeleteHandler(serv, deleteTransport))
	http.Handle("/get_events_for_day", server.NewGetDayHandler(serv, getEventsDayTransport))
	http.Handle("/get_events_for_week", server.NewGetWeekHandler(serv, getEventsWeekTransport))
	http.Handle("/get_events_for_month", server.NewGetMonthHandler(serv, getEventsMonthTransport))
	fmt.Println("starting server at" + port)
	http.ListenAndServe(port, nil)
}
