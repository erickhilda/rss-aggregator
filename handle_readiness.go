package main

import (
	"net/http"
)

func handleReadiness(w http.ResponseWriter, r *http.Request) {
	responseJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}