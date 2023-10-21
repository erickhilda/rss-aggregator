package main

import (
	"log"
	"net/http"
)

func responseErr(w http.ResponseWriter, code int, payload interface{}) {
	if code > 499 {
		log.Printf("Error: %s", payload)
	}

	responseJSON(w, code, payload)
}
