package pkg

import (
	"net/http"
	"testing"
)

var c = Client{
	HostURL:    "http://localhost:8080",
	HTTPClient: &http.Client{},
}
var bookID = 1

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
	book, err := c.GetBook(bookID)
	if err != nil {
		t.Log(err)
	}
	t.Logf("book is %v", book)
}

func TestDeleteBook(t *testing.T) {
	err := c.DeleteBook(bookID)
	if err != nil {
		t.Log(err)
	}
}

func TestUpdateBook(t *testing.T) {
	ub := Book{
		Title:  "About Mostafa",
		Author: "Mostafa",
		Price:  50000,
	}
	err := c.UpdateBook(bookID, ub)
	if err != nil {
		t.Log(err)
	}
}
