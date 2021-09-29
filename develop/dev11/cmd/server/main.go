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

	createTransport := server.NewCreateEventTransport()
	updateTransport := server.NewUpdateEventTransport()
	deleteTransport := server.NewDeleteTransport()

	http.Handle("/createEvent", server.NewCreateHandler(serv, createTransport))
	http.Handle("/updateEvent", server.NewUpdateHandler(serv, updateTransport))
	http.Handle("/deleteEvent", server.NewDeleteHandler(serv, deleteTransport))
	fmt.Println("starting server at: 8081")
	http.ListenAndServe(":8082", nil)
}
