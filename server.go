package main

import (
	"log"
	"net/http"
)

func mustServe(address string, handler http.Handler) {
	var srv = &http.Server{
		Addr:    address,
		Handler: handler,
	}
	log.Printf("Serving endpoint at [%v]\n", srv.Addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("endpoint: %v\n", err)
	}
}
