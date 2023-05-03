package pkg

import (
	"net/http"
	"time"
)

const HostURL string = "localhost:8080"

type Client struct {
	HostURL    string
	HTTPClient *http.Client
}

func NewClient(host *string) (*Client, error) {
	c := Client{
		HostURL: HostURL,
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
	if host != nil {
		c.HostURL = *host
	}
	return &c, nil
}
