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
// }

type myBook struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	ISBN        string `json:"isbn"`
	Description string `json:"description,omitempty"`
}

var books = map[string]myBook{
	"0983721": myBook{Title: "Go For Dummies", Author: "Linkedin Learning", ISBN: "0983721"},
	"3212345": myBook{Title: "Algorithm With Go", Author: "Linkedin Learning", ISBN: "3212345"},
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
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unsupported request method."))
	}
}

func AllBooks() []Book {
	values := make([]Book, len(books))
	idx := 0
	for _, book := range books {
		values[idx] = book
		idx++
	}
}

func writeJSON(w http.ResponseWriter, i interface{}) {
	b, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}
	w.Header().Add("Content-Type", "Application/json; charset:utf-8")
	w.Write(b)
}
