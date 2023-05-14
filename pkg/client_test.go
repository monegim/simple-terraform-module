package pkg

import (
	"net/http"
	"testing"
)

var c = Client{
	HostURL:    "http://localhost:8080",
	HTTPClient: &http.Client{},
}
var id = 1
func TestCreateBook(t *testing.T) {
	b := Book{
		ID:     id,
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
	book, err := c.GetBook(id)
	if err != nil {
		t.Log(err)
	}
	t.Logf("book is %v", book)
}

func TestDeleteBook(t *testing.T) {
	err := c.DeleteBook(id)
	if err != nil {
		t.Log(err)
	}
}

func TestUpdateBook(t *testing.T) {
	ub := Book{
		ID:     id,
		Title:  "About Mostafa",
		Author: "Mostafa",
		Price:  50000,
	}
	err := c.UpdateBook(id, ub)
	if err != nil {
		t.Log(err)
	}
}
