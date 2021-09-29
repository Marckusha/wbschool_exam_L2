package client

import "net/http"

type client struct {
	cli *http.Client
}

