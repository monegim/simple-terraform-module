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
		ID:     1,
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
	id := 1
	book, err := c.GetBook(id)
	if err != nil {
		t.Log(err)
	}
	t.Logf("book is %v", book)
}

func TestDeleteBook(t *testing.T) {
	id := 1
	err := c.DeleteBook(id)
	if err != nil {
		t.Log(err)
	}
}
