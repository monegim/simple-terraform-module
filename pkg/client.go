package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
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

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}
	return body, err
}

func (c *Client) CreateBook(book Book) error {
	b, err := json.Marshal(book)
	if err != nil {
		return err
	}
	_, err = http.NewRequest("POST", fmt.Sprintf("%s/books", c.HostURL), strings.NewReader(string(b)))
	if err != nil {
		return err
	}
	return nil
}
