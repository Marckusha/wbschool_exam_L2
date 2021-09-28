package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func main() {

	conn, err := net.Dial("tcp", "golang.org:80")
	conn.SetDeadline(time.Now().Add(2 * time.Second))
	//conn, err := net.DialTimeout("tcp", "golang.org:80", time.Second)

	if err != nil {
		fmt.Println(err, "error")
		return
	}
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	status, err := bufio.NewReader(conn).ReadString('\n')
	fmt.Println(status)
}
