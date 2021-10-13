package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"wbschool_exam_L2/develop/dev11/pkg/server"

	"github.com/gorilla/mux"
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

	router := mux.NewRouter()

	router.Handle("/create_event", server.NewCreateHandler(serv, createTransport))
	router.Handle("/update_event", server.NewUpdateHandler(serv, updateTransport))
	router.Handle("/delete_event", server.NewDeleteHandler(serv, deleteTransport))
	router.Handle("/get_events_for_day", server.NewGetDayHandler(serv, getEventsDayTransport))
	router.Handle("/get_events_for_week", server.NewGetWeekHandler(serv, getEventsWeekTransport))
	router.Handle("/get_events_for_month", server.NewGetMonthHandler(serv, getEventsMonthTransport))
	http.Handle("/", router)
	fmt.Println("starting server at" + port)
	http.ListenAndServe(port, nil)
}
