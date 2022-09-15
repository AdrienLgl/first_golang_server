package main

import (
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "Hello, world!! \n")
}

func second_route(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "This is the second route !")
}

func router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handler)
	r.HandleFunc("/second-route", second_route)
	return r
}

func main() {
	router := router()
	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:9100",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	println("Server is listening on http://127.0.0.1:9100")
	log.Fatal(srv.ListenAndServe())
}
