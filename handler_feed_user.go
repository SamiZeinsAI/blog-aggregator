package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/SamiZeinsAI/blog-aggregator/internal/database"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerFeedsUserGet(w http.ResponseWriter, r *http.Request, user database.User) {
	feeds, err := cfg.DB.GetFeedsUser(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, 500, "error getting users feeds from database")
		return
	}
	respondWithJSON(w, 200, feeds)
}

func (cfg *apiConfig) handlerFeedUserDelete(w http.ResponseWriter, r *http.Request, user database.User) {
	urlParam := chi.URLParam(r, "feedFollowID")
	feedId, err := uuid.FromBytes([]byte(urlParam))
	if err != nil {
		respondWithError(w, 500, "invalid feed id")
		return
	}
	_, err = cfg.DB.DeleteFeedUser(r.Context(), feedId)
	if err != nil {
		respondWithError(w, 400, "given feed id not in database")
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (cfg *apiConfig) handlerFeedUserCreate(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedId uuid.UUID `json:"feed_id"`
	}
	params := parameters{}
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		respondWithError(w, 400, "error decoding request body")
		return
	}
	usersFeeds, err := cfg.DB.CreateFeedUser(r.Context(), database.CreateFeedUserParams{
		ID:        uuid.New(),
		FeedID:    params.FeedId,
		UserID:    user.ID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		respondWithError(w, 500, "error creating entry in database")
		return
	}
	respondWithJSON(w, 200, databaseUsersFeedsToUsersFeeds(usersFeeds))
}
