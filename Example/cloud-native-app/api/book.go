package api

import (
	"encoding/json"
	"net/http"
)

type myBook struct {
	Title  string
	Author string
	ISBN   string
}

func (b myBook) ToJSON() []byte {
	ToJSON, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	return ToJSON
}

func FromJSON(data []byte) myBook {
	book := myBook{}
	err := json.Unmarshal(data, &book)
	if err != nil {
		panic(err)
	}
	return book
}

var Books = []myBook{
	myBook{Title: "Go For Dummies", Author: "Linkedin Learning", ISBN: "2321231"},
	myBook{Title: "Algorithm With Go", Author: "Coursera", ISBN: "2321231"},
}

func BooksHandleFunc(w http.ResponseWriter, r *http.Request) {
	book, err := json.Marshal(Books)
	if err != nil {
		panic(err)
	}

	w.Header().Add("Content-Type", "application/json; charsetutf-8")
	w.Write(book)
}
