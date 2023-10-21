package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func handleError(w http.ResponseWriter, r *http.Request) {
	responseErr(w, http.StatusInternalServerError, map[string]string{"status": "error", "message": "Something went wrong"})
}

func main() {
	godotenv.Load(".env")
	portstring := os.Getenv("PORT")
	if portstring == "" {
		fmt.Println("PORT not set")
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Content-Length", "Content-Type", "Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/health", handleReadiness)
	v1Router.Get("/error", handleError)

	router.Mount("/v1", v1Router)

	serve := &http.Server{
		Handler: router,
		Addr:    ":" + portstring,
	}

	log.Printf("Listening on port %s", portstring)

	err := serve.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
