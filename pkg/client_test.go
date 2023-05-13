package pkg

import (
	"net/http"
	"testing"
)

var c = Client{
	HostURL:    "http://localhost:8080",
	HTTPClient: &http.Client{},
}

func TestCreateBook(t *testing.T) {
	b := Book{
		Title:  "Atomic Habits",
		Author: "Not Sure",
		Price:  30000,
	}
	err := c.CreateBook(b)
	if err != nil {
		t.Log(err)
	}
}

func TestGetBook(t *testing.T) {
	id := 0
	book, err := c.GetBook(id)
	if err != nil {
		t.Log(err)
	}
	t.Logf("book is %v", book)
}
