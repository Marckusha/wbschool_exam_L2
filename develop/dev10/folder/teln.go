package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

/*
Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080
go-telnet --timeout=3s 1.1.1.1 123
Программа должна подключаться к указанному хосту (ip или
доменное имя) и порту по протоколу TCP После подключения
STDIN программы должен записываться в сокет, а данные
полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение
к серверу (через аргумент --timeout, по умолчанию 10s).
При нажатии Ctrl+D программа должна закрывать сокет и
завершаться. Если сокет закрывается со стороны сервера,
программа должна также завершаться. При подключении к
несуществующему сервер, программа должна завершаться через
timeout.
*/

type TelnetClient struct {
	Address string
	Port    string
	Timeout time.Duration
	reader  *bufio.Reader
	writer  *bufio.Writer
	conn    net.Conn
}

func (tc *TelnetClient) Dial() (err error) {

	//d := net.Dialer{Timeout: tc.Timeout}
	//tc.conn, err = d.Dial("tcp", tc.Address+":"+tc.Port)
	tc.conn, err = net.DialTimeout("tcp", tc.Address+":"+tc.Port, tc.Timeout)
	//tc.conn, err = net.Dial("tcp", tc.Address+":"+tc.Port)
	if err != nil {
		return
	}

	tc.reader = bufio.NewReader(tc.conn)
	tc.writer = bufio.NewWriter(tc.conn)
	//err = tc.conn.SetReadDeadline(time.Now().Add(tc.Timeout))
	if err != nil {
		return
	}

	return
}

func (tc *TelnetClient) Close() {
	tc.conn.Close()
}

func (tc *TelnetClient) Execute() {
	for {
		var source string
		fmt.Print("Введите слово: ")
		_, err := fmt.Scanln(&source)
		if err != nil {
			fmt.Println("Некорректный ввод", err)
			continue
		}
		// отправляем сообщение серверу
		if n, err := tc.conn.Write([]byte(source)); n == 0 || err != nil {
			fmt.Println(err)
			return
		}
		// получем ответ
		fmt.Print("Перевод:")
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

func main() {
	tc := TelnetClient{
		Address: "golang.org",
		Port:    "80",
		Timeout: time.Second,
	}

	err := tc.Dial()
	if err != nil {
		fmt.Println(err)
		return
	}
	tc.Execute()
	defer tc.Close()
}
