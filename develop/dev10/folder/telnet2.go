package main

import (
	"bufio"
	"net"
	"os"
	"time"

	"github.com/reiver/go-oi"
	"github.com/reiver/go-telnet"
)

type MyCaller struct{}

func (c MyCaller) CallTELNET(ctx telnet.Context, w telnet.Writer, r telnet.Reader) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		oi.LongWrite(w, scanner.Bytes())
		oi.LongWrite(w, []byte("\n"))
	}
}

func DialTimeout(addr string, dur time.Duration) (*net.Conn, error) {

	const network = "tcp"

	if "" == addr {
		addr = "127.0.0.1:telnet"
	}

	conn, err := net.DialTimeout(network, addr, dur)
	if nil != err {
		return nil, err
	}

	dataReader := bufio.NewReader(conn)
	dataWriter := bufio.NewWriter(conn)

	return &clientConn, nil
}

func DialTimeoutAndCall(srvAddr string, caller telnet.Caller, dur time.Duration) error {
	conn, err := DialTimeout(srvAddr)
	if nil != err {
		return err
	}

	client := &Client{Caller: caller}

	return client.Call(conn)
}

func main() {

}
