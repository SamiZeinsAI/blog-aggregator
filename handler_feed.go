package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/SamiZeinsAI/blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerFeedCreate(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
	type returnVals struct {
		Id        string `json:"id"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
		Name      string `json:"name"`
		Url       string `json:"url"`
		UserId    string `json:"user_id"`
	}
	params := parameters{}
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		respondWithError(w, 400, "error decoding request body")
		return
	}

	feed, err := cfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,
	})
	if err != nil {
		respondWithError(w, 400, "error creating feed")
		return
	}
	respBody := returnVals{
		Id:        feed.ID.String(),
		CreatedAt: feed.CreatedAt.String(),
		UpdatedAt: feed.UpdatedAt.String(),
		Name:      feed.Name,
		Url:       feed.Url,
		UserId:    feed.UserID.String(),
	}
	respondWithJSON(w, 200, respBody)
}
