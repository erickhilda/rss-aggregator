package main

import (
	"fmt"
	"net/http"

	"github.com/erickhilda/rssagg/auth"
	"github.com/erickhilda/rssagg/internal/db"
)

type authedHandler func(http.ResponseWriter, *http.Request, db.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetApiKey(r.Header)

		if err != nil {
			responseJSON(w, http.StatusUnauthorized, err)
			return
		}
	
		user, err := apiCfg.Db.GetUserByApiKey(r.Context(), apiKey)
		if err != nil {
			responseJSON(w, http.StatusInternalServerError, fmt.Sprintf("error getting user: %s", err))
			return
		}
	
		handler(w, r, user)
	}
}
