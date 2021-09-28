package main

import (
	"fmt"
	"net/http"
	"wbschool_exam_L2/develop/dev11/pkg/server"
)

//check ports
//ss -lntu

func main() {
	serv := server.NewService()

	tranportCreateEvent := server.NewCreateEventTransport()
	transportUpdateEven := server.NewUpdateEventTransport()

	http.Handle("/createEvent", server.NewCreateHandler(serv, tranportCreateEvent))
	http.Handle("/updateEvent", server.NewUpdateHandler(serv, transportUpdateEven))
	fmt.Println("starting server at: 8082")
	http.ListenAndServe(":8082", nil)
}
