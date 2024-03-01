package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/SamiZeinsAI/blog-aggregator/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	port string
	DB   *database.Queries
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	dbUrl := os.Getenv("POSTGRES")
	conn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal(err)
	}

	apiCfg := apiConfig{
		port: os.Getenv("PORT"),
		DB:   database.New(conn),
	}

	go apiCfg.scraper()

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	v1Router := chi.NewRouter()

	v1Router.Get("/readiness", apiCfg.handlerReadiness)
	v1Router.Get("/err", apiCfg.handlerErr)

	v1Router.Post("/users", apiCfg.handlerUsersCreate)
	v1Router.Get("/users", apiCfg.middlewareAuth(apiCfg.handlerUsersGet))

	v1Router.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handlerFeedCreate))
	v1Router.Get("/feeds", apiCfg.handlerFeedGetAll)

	v1Router.Post("/feed_follows", apiCfg.middlewareAuth(apiCfg.handlerFeedUserCreate))
	v1Router.Delete("/feed_follows/{feedFollowID}", apiCfg.middlewareAuth(apiCfg.handlerFeedUserDelete))
	v1Router.Get("/feed_follows", apiCfg.middlewareAuth(apiCfg.handlerFeedsUserGet))

	router.Mount("/v1", v1Router)
	srv := &http.Server{
		Addr:    ":" + apiCfg.port,
		Handler: router,
	}

	log.Printf("Serving on port: %s\n", apiCfg.port)
	log.Fatal(srv.ListenAndServe())
}
