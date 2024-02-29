package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/SamiZeinsAI/blog-aggregator/internal/database"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerFeedFollowsGet(w http.ResponseWriter, r *http.Request, user database.User) {
	feeds, err := cfg.DB.GetUsersFeedFollows(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, 500, "error getting users feed follows from database")
		return
	}
	respondWithJSON(w, 200, feeds)
}

func (cfg *apiConfig) handlerFeedFollowDelete(w http.ResponseWriter, r *http.Request, user database.User) {
	urlParam := chi.URLParam(r, "feedFollowID")
	feedId, err := uuid.FromBytes([]byte(urlParam))
	if err != nil {
		respondWithError(w, 500, "invalid feed id")
		return
	}
	_, err = cfg.DB.DeleteFeedFollow(r.Context(), feedId)
	if err != nil {
		respondWithError(w, 400, "given feed id not in database")
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (cfg *apiConfig) handlerFeedFollowCreate(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedId uuid.UUID `json:"feed_id"`
	}
	type returnVals struct {
		Id        string `json:"id"`
		FeedId    string `json:"feed_id"`
		UserId    string `json:"user_id"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}
	params := parameters{}
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		respondWithError(w, 400, "error decoding request body")
		return
	}
	feedFollow, err := cfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
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
	respBody := returnVals{
		Id:        feedFollow.ID.String(),
		FeedId:    feedFollow.FeedID.String(),
		UserId:    feedFollow.UserID.String(),
		CreatedAt: feedFollow.CreatedAt.String(),
		UpdatedAt: feedFollow.UpdatedAt.String(),
	}
	respondWithJSON(w, 200, respBody)
}
