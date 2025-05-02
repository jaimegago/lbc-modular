package main

import (
	"net/http"
	// "github.com/jaimegago/lbc-modular/fizzbuzz"
)

type fbHandler struct{}

func (h *fbHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.RequestURI == "/fizzbuzz" && r.Method == http.MethodPost:
		h.GetFbResults(w, r)
	case r.RequestURI == "/fizzbuzz-stats":
		h.GetFbStats(w, r)
	default:
		h.Default(w, r)
	}
}

func (h *fbHandler) GetFbResults(w http.ResponseWriter, r *http.Request) {
	// fb := fizzbuzz.fizzbuzz{}
	// log.Println(fb)
	w.Write([]byte("got a var"))
}

func (h *fbHandler) GetFbStats(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this is fb stats"))
}

func (h *fbHandler) Default(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this should be the docs page"))
}

func InternalServerErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("500 Internal Server Error"))
}

func main() {

	mux := http.NewServeMux()

	mux.Handle("/fizzbuzz", &fbHandler{})
	mux.Handle("/fizzbuzz-stats", &fbHandler{})

	http.ListenAndServe(":8080", mux)
}
