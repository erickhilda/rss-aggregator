package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/erickhilda/rssagg/internal/db"
	"github.com/go-chi/chi"
)

func (apiCfg *apiConfig) handleCreateFeedFollow(w http.ResponseWriter, r *http.Request, user db.User) {
	type parameters struct {
		FeedID int32 `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)

	params := &parameters{}
	err := decoder.Decode(params)
	if err != nil {
		responseJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request"})
		return
	}

	feedId := uint64(params.FeedID)
	_, err = apiCfg.Db.GetFeedByID(r.Context(), feedId)
	if err != nil {
		responseJSON(w, http.StatusBadRequest, map[string]string{"error": "cannot find feed"})
		return
	}

	createFeedFollowParams := db.CreateFeedFollowParams{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		FeedID:    params.FeedID,
		UserID:    user.ID,
	}

	feedFollows, err := apiCfg.Db.CreateFeedFollow(r.Context(), createFeedFollowParams)
	if err != nil {
		responseJSON(w, http.StatusInternalServerError, err)
		return
	}
	feedResult, errInserted := feedFollows.LastInsertId()
	if errInserted != nil {
		responseJSON(w, http.StatusInternalServerError, errInserted)
		return
	}

	responseJSON(w, http.StatusOK, map[string]int64{"id": feedResult})
}

func (apiCfg *apiConfig) handleGetFeedFollowsByUserId(w http.ResponseWriter, r *http.Request, user db.User) {
	feedFollows, err := apiCfg.Db.GetFeedFollowsByUserID(r.Context(), user.ID)
	if err != nil {
		responseJSON(w, http.StatusInternalServerError, err)
		return
	}

	responseJSON(w, http.StatusOK, databaseFeedFollowsToFeedFollows(feedFollows))
}

func (apiCfg *apiConfig) handleDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user db.User) {
	feedFollowId, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		responseJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request"})
		return
	}

	_, err = apiCfg.Db.GetFeedFollowByID(r.Context(), feedFollowId)
	if err != nil {
		responseJSON(w, http.StatusBadRequest, map[string]string{"error": "cannot find feed follow"})
		return
	}

	deleteFeedFollowParams := db.DeleteFeedFollowParams{
		ID: feedFollowId,
		UserID: user.ID,
	}
	err = apiCfg.Db.DeleteFeedFollow(r.Context(), deleteFeedFollowParams)
	if err != nil {
		responseJSON(w, http.StatusInternalServerError, err)
		return
	}

	responseJSON(w, http.StatusOK, map[string]string{"status": "deleted"})
}
