package main

import "net/http"

func (cfg *apiConfig) GetErrorHandler(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, 500, "Internal Sever Error")
}
