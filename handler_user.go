package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/erickhilda/rssagg/internal/db"
	"github.com/google/uuid"
)

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

	createUserParams := db.CreateUserParams{
		ID:        uuid.New().String(),
		Email:     params.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
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

	responseJSON(w, http.StatusOK, databaseUserToUser(createUserParams))
}
