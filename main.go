package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	err := http.ListenAndServe(":5050", http.HandlerFunc(mux))
	log.Println(err)
}

func mux(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	switch r.URL.Path {
	case "/":
		indexHandler(w, r)
	case "/about":
		aboutHandler(w, r)
	default:
		http.NotFound(w, r)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Index Page"))
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About Page"))
}
func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("404 Page Not Found"))
}
