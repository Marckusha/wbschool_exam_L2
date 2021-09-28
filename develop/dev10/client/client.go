package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"
	"time"
)

type TelnetClient struct {
	Address string
	Port    string
	Timeout time.Duration
	reader  *bufio.Reader
	writer  *bufio.Writer
	conn    net.Conn
}

func (tc *TelnetClient) Dial() (err error) {

	if tc.Timeout == 0 {
		tc.Timeout = 10 * time.Second
	}

	tc.conn, err = net.DialTimeout("tcp", tc.Address+":"+tc.Port, tc.Timeout)
	if err != nil {
		return
	}

	//tc.reader = bufio.NewReader(tc.conn)
	//tc.writer = bufio.NewWriter(tc.conn)

	return
}

func (tc *TelnetClient) Close() {
	tc.conn.Close()
}

func (tc *TelnetClient) Execute() {
	for {

		var text string
		_, err := fmt.Scanln(&text)

		if err == io.EOF {
			fmt.Println("EOF")
		}

		if err != nil {
			fmt.Println("Некорректный ввод", err)
			continue
		}
		// отправляем сообщение серверу
		if n, err := tc.conn.Write([]byte(text)); n == 0 || err != nil {
			fmt.Println(err)
			return
		}

		// получем ответ
		buff := make([]byte, 1024)
		n, err := tc.conn.Read(buff)
		if err != nil {
			fmt.Println("TUT", err)
			break
		}
		fmt.Print(string(buff[0:n]))
		fmt.Println()
	}

}

var (
	timeout = flag.String("timeout", "10s", "UDP")
)

func StrTimeToInt(str, suf string) (int, error) {
	val, err := strconv.Atoi(strings.TrimSuffix(str, suf))
	if err != nil {
		return -1, err
	}
	return val, nil
}

func StringToTime(t string) (dur time.Duration) {

	if strings.HasSuffix(t, "ms") {
		val, _ := StrTimeToInt(t, "ms")
		dur = time.Millisecond * time.Duration(val)
	} else if strings.HasSuffix(t, "ns") {
		val, _ := StrTimeToInt(t, "ns")
		dur = time.Nanosecond * time.Duration(val)
	} else if strings.HasSuffix(t, "mu") {
		val, _ := StrTimeToInt(t, "mu")
		dur = time.Microsecond * time.Duration(val)
	} else if strings.HasSuffix(t, "s") {
		val, _ := StrTimeToInt(t, "s")
		dur = time.Second * time.Duration(val)
	} else if strings.HasSuffix(t, "m") {
		val, _ := StrTimeToInt(t, "m")
		dur = time.Second * time.Duration(val)
	} else if strings.HasSuffix(t, "h") {
		val, _ := StrTimeToInt(t, "h")
		dur = time.Hour * time.Duration(val)
	}

	return
}

func main() {
	flag.Parse()

	tc := TelnetClient{
		Address: "localhost",
		Port:    "8084",
		Timeout: StringToTime(*timeout),
	}

	err := tc.Dial()
	if err != nil {
		fmt.Println(err)
		return
	}
	tc.Execute()

	defer tc.Close()
}
