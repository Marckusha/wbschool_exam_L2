package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

/*
type Event struct {
	Day   string
	Week  string
	Month string
	Task  string
}*/

func main() {

	s := []byte(`{"Day":"1","Month":"1","Week":"may","Task":"print hello world"}`)
	r := bytes.NewReader(s)
	cli := http.Client{}
	fmt.Println("1")
	resp, err := cli.Post("http://localhost:8082/deleteEvent", "application/json", r)
	fmt.Println("2")
	if err != nil {
		fmt.Println("error")
		return
	}
	fmt.Println("3")
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
	fmt.Println("4")
}
