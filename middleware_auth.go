package main

import (
	"net/http"

	"github.com/SamiZeinsAI/blog-aggregator/internal/auth"
	"github.com/SamiZeinsAI/blog-aggregator/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 400, "error getting api key")
			return
		}
		user, err := cfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 400, "error getting user")
			return
		}
		handler(w, r, user)
	}
}
