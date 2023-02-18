package http

import (
	"net/http"
	"time"
)

func Client() http.Client {
	client := http.Client{
		Timeout: time.Second,
	}

	return client
}
