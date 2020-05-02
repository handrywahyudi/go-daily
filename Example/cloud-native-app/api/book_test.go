package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBookToJson(t *testing.T) {
	book := myBook{Title: "Abrakadabra", Author: "Brama Kumbara", ISBN: "085232143"}
	json := book.ToJSON()

	assert.Equal(t, `{"Title":"Abrakadabra","Author":"Brama Kumbara","ISBN":"085232143"}`,
		string(json), "Json Marshalling Wrong.")
}

func TestBookFromJson(t *testing.T) {
	json := []byte(`{"Title":"Abrakadabra","Author":"Brama Kumbara","ISBN":"085232143"}`)
	book := FromJSON(json)

	assert.Equal(t, myBook{Title: "Abrakadabra", Author: "Brama Kumbara", ISBN: "085232143"},
		book, "Json Unmarshalling Wrong.")
}
