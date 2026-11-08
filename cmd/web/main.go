package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home2)
	mux.HandleFunc("/snippet/view", snippetView2)
	mux.HandleFunc("/snippet/create", snippetCreate2)

	log.Print("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)

	log.Fatal("error: ", err)

	/*mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	log.Print("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)*/
}

func home2(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello from Snippetbox"))
}

func snippetView2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet..."))
}

func snippetCreate2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new snippet"))
}
