package config

import (
	"database/sql"
	"log"

	"github.com/brunocapri/movie-engine/internal/assets"
	"github.com/brunocapri/movie-engine/internal/cache"
	"github.com/brunocapri/movie-engine/openai"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

type Application struct {
	OpenAiClient *openai.Client
	Cache        *cache.Cache
}

func InitMux() *chi.Mux {
	mux := chi.NewRouter()
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	assets.StartStaticServer(mux)

	return mux
}

func InitDb(dbURL string) *sql.DB {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
