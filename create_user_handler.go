package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/SamiZeinsAI/blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) PostUserHandler(w http.ResponseWriter, r *http.Request) {
	type paramters struct {
		Name string `json:"name"`
	}
	type returnVals struct {
		Id        string `json:"id"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
		Name      string `json:"name"`
	}
	params := paramters{}
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		respondWithError(w, 400, "Error decoding request body")
		return
	}
	fmt.Printf("%s\n", params.Name)

	user, err := cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		respondWithError(w, 500, "Error posting user to database")
		return
	}
	respBody := returnVals{
		Id:        user.ID.String(),
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
		Name:      params.Name,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(respBody)
	if err != nil {
		respondWithError(w, 500, "Error encoding response body")
		return
	}
}
