package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	api "github.com/mesg-foundation/core/api/service"
)

type simpleEndpoint struct {
	now       func() string
	uuid      func() string
	emitEvent func(body string) (*api.EmitEventReply, error)
}

type eventOnRequest struct {
	Date string `json:"date"`
	ID   string `json:"id"`
	Body string `json:"body"`
}

func (e simpleEndpoint) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	var body, err = ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("endpoint: failed to read payload: %v\n", err)
		http.Error(w, "Invalid payload.", http.StatusBadRequest)
		return
	}
	if err := r.Body.Close(); err != nil {
		log.Printf("endpoint: failed to close playload reader: %v\n", err)
		http.Error(w, "Failed to read payload.", http.StatusInternalServerError)
		return
	}

	var ev, _ = json.Marshal(eventOnRequest{
		Date: e.now(),
		ID:   e.uuid(),
		Body: string(body),
	})

	reply, err := e.emitEvent(string(ev))
	if err != nil {
		log.Printf("endpoint: failed to emit event: %v\n", err)
		http.Error(w, "Failed to process request.", http.StatusInternalServerError)
		return
	}

	log.Printf("endpoint: reply from mesg: %v\n", reply)
	w.WriteHeader(http.StatusOK)
}
