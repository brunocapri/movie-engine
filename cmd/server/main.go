package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/brunocapri/movie-engine/internal/config"
	"github.com/brunocapri/movie-engine/internal/handlers"
	"github.com/brunocapri/movie-engine/internal/repository"
	"github.com/brunocapri/movie-engine/openai"
	_ "github.com/lib/pq"
)

func main() {
	cfg := config.InitConfig()
	mux := config.InitMux()
	db := config.InitDb(cfg.DB())
	openAiClient := openai.NewClient(cfg.OpenAIUri(), cfg.OpenAIKey())

	app := &config.Application{
		Config:       cfg,
		Mux:          mux,
		Db:           db,
		OpenAiClient: openAiClient,
	}

	movies := repository.NewMovieRepository(db)

	handlers := handlers.NewHandler(app, movies)
	mux.Get("/", handlers.GetHome)
	mux.Get("/search", handlers.GetSearch)

	server := &http.Server{
		Addr:    ":" + cfg.Port(),
		Handler: http.TimeoutHandler(mux, 30*time.Second, "request timed out"),
	}

	fmt.Printf("🚂 Movie Engine running on: http://localhost:%s\n", cfg.Port())

	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
