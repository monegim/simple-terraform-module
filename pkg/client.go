package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
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

	body, err := ioutil.ReadAll(res.Body)
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
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/books", c.HostURL), strings.NewReader(string(b)))
	if err != nil {
		return err
	}
	_, err = c.doRequest(req)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) GetBook(bookID int) (*Book, error) {
	book := Book{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/books/%s", c.HostURL, strconv.Itoa(bookID)), nil)
	if err != nil {
		return nil, err
	}
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &book)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (c *Client) UpdateBook(bookId int, updatedBook Book) error {
	b, err := json.Marshal(updatedBook)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/books/%s", c.HostURL, strconv.Itoa(bookId)), strings.NewReader(string(b)))
	if err != nil {
		return err
	}
	_, err = c.doRequest(req)
	if err != nil {
		return err
	}
	return nil
}
func (c *Client) DeleteBook(bookID int) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/books/%s", c.HostURL, strconv.Itoa(bookID)), nil)
	if err != nil {
		return err
	}
	_, err = c.doRequest(req)
	if err != nil {
		return nil
	}
	return nil
}
