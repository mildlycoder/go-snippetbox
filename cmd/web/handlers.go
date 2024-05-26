package main

import (
	"fmt"
	"net/http"
	"strconv"
    "html/template"
    "log"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
    files := []string{
        "./ui/html/pages/home.html", 
        "./ui/html/base.html",
    "./ui/html/partials/nav.html"}
    ts, err := template.ParseFiles(files...)
    if err != nil {
        log.Println(err.Error())
        http.Error(w, "Internal server error",http.StatusInternalServerError)
        return
    }
    err = ts.ExecuteTemplate(w, "base", nil)
    if err != nil {
        log.Printf(err.Error())
        http.Error(w, "Internal server error", http.StatusInternalServerError)
    }
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Displaying snippet with Id: %d", id)
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Create snippet from here"))
}
