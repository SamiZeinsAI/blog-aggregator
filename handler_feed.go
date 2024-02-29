package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/SamiZeinsAI/blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerFeedGetAll(w http.ResponseWriter, r *http.Request) {
	feeds, err := cfg.DB.GetAllFeeds(r.Context())
	if err != nil {
		respondWithError(w, 500, "error getting feeds")
		return
	}
	respondWithJSON(w, 200, feeds)
}

func (cfg *apiConfig) handlerFeedCreate(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
	type returnVals struct {
		Feed       Feed      `json:"feed"`
		FeedFollow FeedsUser `json:"feed_follow"`
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
		respondWithError(w, 500, "error creating feed")
		return
	}
	feedUser, err := cfg.DB.CreateFeedUser(r.Context(), database.CreateFeedUserParams{
		ID:        uuid.New(),
		FeedID:    feed.ID,
		UserID:    user.ID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		respondWithError(w, 500, "error creating feed follow")
		return
	}
	respondWithJSON(w, 200, returnVals{
		Feed:       databaseFeedToFeed(feed),
		FeedFollow: databaseUsersFeedsToUsersFeeds(feedUser),
	})
}
