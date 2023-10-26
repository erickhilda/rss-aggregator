package main

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"time"

	"github.com/erickhilda/rssagg/internal/db"
	"github.com/google/uuid"
)

func generateAPIKey(length int) (string, error) {
	// Create a byte slice to store random bytes
	keyBytes := make([]byte, length/2) // We divide by 2 since each byte is represented by 2 hexadecimal characters

	// Read random bytes into the byte slice
	_, err := rand.Read(keyBytes)
	if err != nil {
		return "", err
	}

	// Convert the random bytes to a hexadecimal string
	apiKey := hex.EncodeToString(keyBytes)

	return apiKey, nil
}

func (apiCfg *apiConfig) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Email string `json:"email"`
	}

	decoder := json.NewDecoder(r.Body)

	params := &parameters{}
	err := decoder.Decode(params)
	if err != nil {
		responseJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request"})
		return
	}

	apiKey, err := generateAPIKey(64)
	if err != nil {
		responseJSON(w, http.StatusInternalServerError, err)
		return
	}

	createUserParams := db.CreateUserParams{
		ID:        uuid.New().String(),
		Email:     params.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		ApiKey:    apiKey,
	}

	user, err := apiCfg.Db.CreateUser(r.Context(), createUserParams)
	if err != nil {
		responseJSON(w, http.StatusInternalServerError, err)
		return
	}
	_, errInserted := user.LastInsertId()
	if errInserted != nil {
		responseJSON(w, http.StatusInternalServerError, errInserted)
		return
	}

	createdUser := db.User{
		ID:        createUserParams.ID,
		Email:     createUserParams.Email,
		CreatedAt: createUserParams.CreatedAt,
		UpdatedAt: createUserParams.UpdatedAt,
		ApiKey:    createUserParams.ApiKey,
	}

	responseJSON(w, http.StatusOK, databaseUserToUser(createdUser))
}

func (apiCfg *apiConfig) handleGetUser(w http.ResponseWriter, r *http.Request, user db.User) {
	responseJSON(w, http.StatusOK, databaseUserToUser(user))
}

func (apiCfg *apiConfig) handleGetPostsForUser(w http.ResponseWriter, r *http.Request, user db.User) {
	posts, err := apiCfg.Db.GetPostsForUser(r.Context(), db.GetPostsForUserParams{
		UserID: user.ID,
		Limit: 10,
	})
	if err != nil {
		responseJSON(w, http.StatusInternalServerError, err)
		return
	}

	responseJSON(w, http.StatusOK, databasePostsToPosts(posts))
}
