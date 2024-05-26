package main

import (
	"log"
	"net/http"
)

func main() {
    fileserver := http.FileServer(http.Dir("./ui/static/"))
    http.Handle("/static/", http.StripPrefix("/static/", fileserver))
	http.HandleFunc("/", home)
	http.HandleFunc("/snippet/view", snippetView)
	http.HandleFunc("/snippet/create", createSnippet)
	log.Println("Starting server  on :4000")
	err := http.ListenAndServe(":4000", nil)
	log.Fatal(err)
}
