package pkg

import (
	"net/http"
	"testing"
)

func TestCreateBook(t *testing.T) {
	b := Book{
		Title:  "Atomic Habits",
		Author: "Not Sure",
		Price:  30000,
	}
	c := Client{
		HostURL:    "http://localhost:8080",
		HTTPClient: &http.Client{},
	}
	err := c.CreateBook(b)
	if err != nil {
		t.Log(err)
	}
}
