package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

type apiConfig struct {
	port string
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	apiCfg := apiConfig{
		port: os.Getenv("PORT"),
	}

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{}))
	v1Router := chi.NewRouter()

	v1Router.Get("/readiness", apiCfg.GetReadinessHandler)
	v1Router.Get("/err", apiCfg.GetErrorHandler)

	router.Mount("/v1", v1Router)
	srv := &http.Server{
		Addr:    ":" + apiCfg.port,
		Handler: router,
	}

	log.Printf("Serving on port: %s\n", apiCfg.port)
	log.Fatal(srv.ListenAndServe())
}
