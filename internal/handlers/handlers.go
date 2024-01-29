package handlers

import (
	"github.com/brunocapri/movie-engine/internal/config"
	"github.com/brunocapri/movie-engine/internal/repository"
)

type Handler struct {
	app    *config.Application
	movies *repository.MovieRepository
}

func NewHandler(app *config.Application, movies *repository.MovieRepository) *Handler {
	return &Handler{
		app:    app,
		movies: movies,
	}
}
