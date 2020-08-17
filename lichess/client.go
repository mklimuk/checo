package lichess

import (
	"net/http"
	"time"
)

type Client struct {
	http.Client
	token string
	baseURL string
}

func NewClient(token, baseURL string) *Client {
	return &Client{
		Client: http.Client{Timeout:10 * time.Second},
		token: token,
		baseURL: baseURL,
	}
}


