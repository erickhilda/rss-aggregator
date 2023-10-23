package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/erickhilda/rssagg/internal/db"
)

func (apiCfg *apiConfig) handleCreateFeed(w http.ResponseWriter, r *http.Request, user db.User) {
	type parameters struct {
		Name   string `json:"name"`
		Url    string `json:"url"`
		UserID string `json:"user_id"`
	}

	decoder := json.NewDecoder(r.Body)

	params := &parameters{}
	err := decoder.Decode(params)
	if err != nil {
		responseJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request"})
		return
	}

	createFeedParams := db.CreateFeedParams{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,
	}

	feed, err := apiCfg.Db.CreateFeed(r.Context(), createFeedParams)
	if err != nil {
		responseJSON(w, http.StatusInternalServerError, err)
		return
	}
	feedResult, errInserted := feed.LastInsertId()
	if errInserted != nil {
		responseJSON(w, http.StatusInternalServerError, errInserted)
		return
	}

	responseJSON(w, http.StatusOK, map[string]int64{"id": feedResult})
}


func (apiCfg *apiConfig) handleGetFeed(w http.ResponseWriter, r *http.Request) {

	feeds, err := apiCfg.Db.GetFeed(r.Context())
	if err != nil {
		responseJSON(w, http.StatusInternalServerError, err)
		return
	}

	responseJSON(w, http.StatusOK, databaseFeedsToFeeds(feeds))
}