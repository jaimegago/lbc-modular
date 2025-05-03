package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	fb "github.com/jaimegago/lbc-modular/pkg/fizzbuzz"
)

type fbHandler struct{}

func (h *fbHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.RequestURI == "/fizzbuzz" && r.Method == http.MethodPost:
		h.CreateFbResults(w, r)
	case r.RequestURI == "/fizzbuzz-stats":
		h.GetFbStats(w, r)
	default:
		h.Default(w, r)
	}
}

func (h *fbHandler) CreateFbResults(w http.ResponseWriter, r *http.Request) {
	var fbReq fb.ReqData
	// TODO move json decode to its own func
	err := json.NewDecoder(r.Body).Decode(&fbReq)
	if err != nil {
		log.Println("ERROR: Failed to decode json", err)
		InternalServerErrorHandler(w, r)
		return
	}
	err = fbReq.Validate()
	if err != nil {
		InternalServerErrorHandler(w, r)
		return
	}
	// Get stats from storage
	var stats fb.Stats
	stats.ReqParamsHits = []fb.ReqData{}
	stats.CountReqParamsHit(fbReq)
	log.Println("INFO: stats after count: ", stats)
	err = fbReq.Create()

	if err != nil {
		log.Println("ERROR: failed to create results", err)
		InternalServerErrorHandler(w, r)
		return
	}
	sendJson(w, fbReq.Results)
}

func (h *fbHandler) GetFbStats(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this is fb stats"))
}

func (h *fbHandler) Default(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this should be the docs page"))
}

func InternalServerErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(r.RequestURI + " :500 Internal Server Error"))
}

func sendJson(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w,
			"Internal Server Error while trying to encode/send json: %s\n, contact admin@fizzbuzz.com",
			err.Error())
		return
	}
}

func main() {

	mux := http.NewServeMux()

	mux.Handle("/fizzbuzz", &fbHandler{})
	mux.Handle("/fizzbuzz-stats", &fbHandler{})

	http.ListenAndServe(":8080", mux)
}
