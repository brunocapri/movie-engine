package core

import (
	"fmt"

	"github.com/brunocapri/movie-engine/internal/config"
	"github.com/brunocapri/movie-engine/internal/domain"
	"github.com/brunocapri/movie-engine/internal/repository"
	"github.com/brunocapri/movie-engine/internal/utils"
)

type Core struct {
	app    *config.Application
	movies *repository.MovieRepository
}

func NewCore(app *config.Application, movies *repository.MovieRepository) *Core {
	return &Core{
		app:    app,
		movies: movies,
	}
}

func (c Core) SearchMoviesByEmbedding(query string) ([]domain.Movie, error) {
	movies, hit := c.app.Cache.Get(query)
	if hit {
		return movies, nil
	}

	if movie := c.checkExactMatch(query); movie != nil {
		return c.GetSimilarMovies(*movie, false)
	}

	res, err := c.app.OpenAiClient.GenerateEmbeddings(query)
	if err != nil {
		return nil, err
	}

	embedding := c.app.OpenAiClient.EmbeddingToString(res.Data[0].Embedding)
	movies, err = c.movies.FindByEmbedding(embedding)

	if err != nil {
		return nil, err
	}

	c.app.Cache.Add(query, movies)

	return movies, err
}

func (c Core) GetSimilarMovies(movie domain.Movie, search bool) (movies []domain.Movie, err error) {
	if movies, hit := c.app.Cache.Get(movie.Title); hit {
		return movies, nil
	}

	if search {
		movie, _ = c.movies.FindById(movie.Id)
	}

	movies, err = c.movies.FindSimilar(movie)
	if err != nil {
		fmt.Println(err)
		return
	}
	movies = utils.Prepend[domain.Movie](movies, movie)
	c.app.Cache.Add(movie.Title, movies)
	return
}

func (c Core) checkExactMatch(query string) *domain.Movie {
	movie, err := c.movies.FindByTitle(query)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &movie
}
