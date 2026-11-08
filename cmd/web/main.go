package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home2)

	srv := &http.Server{
		Addr:    ":4000",
		Handler: mux,
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go func() {
		<-ctx.Done()

		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		srv.Shutdown(shutdownCtx)
	}()

	log.Print("Starting server on :4000")
	err := srv.ListenAndServe()

	if errors.Is(err, http.ErrServerClosed) {
		log.Fatal("error: ", err)
	}

	log.Print("server stopped gracefully")

	/*log.Print("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)

	log.Fatal("error: ", err)*/

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
	w.Write([]byte("Hello"))
}
