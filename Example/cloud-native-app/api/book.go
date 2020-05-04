package api

import (
	"encoding/json"
	"net/http"
)

// type myBook struct {
// 	Title  string
// 	Author string
// 	ISBN   string
// }
// var Books = []myBook{
// 	myBook{Title: "Go For Dummies", Author: "Linkedin Learning", ISBN: "2321231"},
// 	myBook{Title: "Algorithm With Go", Author: "Coursera", ISBN: "2321231"},
}

type myBook struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	ISBN        string `json:"isbn"`
	Description string `json:"description,omitempty"`
}

var Books = map[string]Book {
    "0983721": Books{Title:"Go For Dummies", Authro: "Linkedin Learning", ISBN: "0983721"}
    "3212345": Books{Title:"Algorithm With Go", Authro: "Linkedin Learning", ISBN: "3212345"}
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

func BooksHandleFunc(w http.ResponseWriter, r *http.Request) {
	// book, err := json.Marshal(Books)
	// if err != nil {
	// 	panic(err)
	// }

	// w.Header().Add("Content-Type", "application/json; charsetutf-8")
	// w.Write(book)

	switch method := r.Method; method {
	case http.MethodGet:
		books := AllBooks()
		writeJSON(w, books)
	}
}
