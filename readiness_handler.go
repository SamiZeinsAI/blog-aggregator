package main

import "net/http"

func (cfg *apiConfig) GetReadinessHandler(w http.ResponseWriter, r *http.Request) {
	type returnVals struct {
		Status string `json:"string"`
	}
	respBody := returnVals{
		Status: "ok",
	}
	respondWithJSON(w, 200, respBody)
}
