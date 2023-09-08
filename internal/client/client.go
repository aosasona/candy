package client

import "net/http"

type Candy struct {
	HttpClient *http.Client
}

var Client *Candy
