package main

import (
	"fmt"
	"net/http"
        "os"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/cloud", cloud)
	http.HandleFunc("/api/echo", echo)
	http.ListenAndServe(port(), nil)
}

func port() string{
    port := os.Getenv("PORT")
    if len(port) == 0 {
        port = "9090"
    }
    return ":" + port
}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Cloud native app example.")
}

func cloud(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "This cloud page.")
}

func echo(w http.ResponseWriter, r *http.Request) {
    message := r.URL.Query()["message"][0]

    w.Header().Add("Content-type", "text/plain")
    fmt.Fprintf(w, message)
}
