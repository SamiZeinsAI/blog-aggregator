package main

import (
	"net/http"
	"strconv"

	"github.com/SamiZeinsAI/blog-aggregator/internal/database"
	"github.com/go-chi/chi"
)

func (cfg *apiConfig) handlerPostsGetByUser(w http.ResponseWriter, r *http.Request, user database.User) {
	limitStr := chi.URLParam(r, "limit")
	limit := 10
	if limitInt, err := strconv.Atoi(limitStr); err == nil && limit > 0 {
		limit = limitInt
	}
	posts, err := cfg.DB.GetPostsByUser(r.Context(), database.GetPostsByUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		respondWithError(w, 500, "error getting users' posts")
		return
	}
	respondWithJSON(w, 200, posts)
}
