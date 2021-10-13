package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

var port string = ":8083"

func postCreate() {

	s := []byte(`{"Date":"2021-01-01T00:00:00Z","Note": "Happy burdfs"}`)

	r := bytes.NewReader(s)
	cli := http.Client{}

	req, err := http.NewRequest("POST", "http://localhost"+port+"/create_event", r)
	req.Header.Set("user_id", "2")
	resp, err := cli.Do(req)

	if err != nil {
		fmt.Println("error")
		return
	}

	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}

func postDelete() {

	s := []byte(`{"Date":"2021-01-01","Note": "Hello, World!"}`)
	r := bytes.NewReader(s)
	cli := http.Client{}
	fmt.Println("1")
	req, err := http.NewRequest("POST", "http://localhost"+port+"/delete_event", r)
	req.Header.Set("user_id", "2")
	resp, err := cli.Do(req)
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

func get() {
	url := "http://localhost" + port + "/get_events_for_month?user_id=1&data=2021-01-01"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("dont get request")
		return
	}

	defer resp.Body.Close()
	fmt.Println("end get method")
}

func main() {
	postCreate()
}
