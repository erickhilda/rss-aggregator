package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/erickhilda/rssagg/internal/db"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type apiConfig struct {
	Db *db.Queries
}

func main() {
	godotenv.Load(".env")

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		fmt.Println("DB_URL not set")
	}

	conn, err := sql.Open("mysql", dbUrl)
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	apiCfg := &apiConfig{
		Db: db.New(conn),
	}

	go startScapper(apiCfg.Db, 4, time.Minute)

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

	v1Router.Post("/users", apiCfg.handleCreateUser)
	v1Router.Get("/users", apiCfg.middlewareAuth(apiCfg.handleGetUser))

	v1Router.Get("/posts", apiCfg.middlewareAuth(apiCfg.handleGetPostsForUser))
	
	v1Router.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handleCreateFeed))
	v1Router.Get("/feeds", apiCfg.handleGetFeed)
	v1Router.Get("/feeds/{id}", apiCfg.handleGetFeedById)
	
	v1Router.Post("/feed_follows", apiCfg.middlewareAuth(apiCfg.handleCreateFeedFollow))
	v1Router.Get("/feed_follows", apiCfg.middlewareAuth(apiCfg.handleGetFeedFollowsByUserId))
	v1Router.Delete("/feed_follows/{id}", apiCfg.middlewareAuth(apiCfg.handleDeleteFeedFollow))

	router.Mount("/v1", v1Router)

	serve := &http.Server{
		Handler: router,
		Addr:    ":" + portstring,
	}

	log.Printf("Listening on port %s", portstring)

	errHttpServe := serve.ListenAndServe()
	if errHttpServe != nil {
		log.Fatal(errHttpServe)
	}
}
