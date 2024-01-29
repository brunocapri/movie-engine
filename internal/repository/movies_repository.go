package repository

import (
	"database/sql"
	"log"

	"github.com/brunocapri/movie-engine/internal/domain"
)

type MovieRepository struct {
	db *sql.DB
}

func NewMovieRepository(db *sql.DB) *MovieRepository {
	return &MovieRepository{db: db}
}

func (r MovieRepository) FindByEmbedding(embedding string) ([]domain.Movie, error) {
	const query = "SELECT title, year, runtime, director, poster, ratings FROM omdb_movies ORDER BY embedding <-> $1 LIMIT 5"
	rows, err := r.db.Query(query, embedding)
	if err != nil {
		log.Printf("Query error: %v", err)
		return nil, err
	}
	defer rows.Close()

	movies := []domain.Movie{}
	for rows.Next() {
		movie := domain.Movie{}
		err := rows.Scan(&movie.Title, &movie.Year, &movie.Runtime, &movie.Director, &movie.Poster, &movie.Ratings)
		if err != nil {
			log.Printf("Scan error: %v", err)
			continue
		}
		movies = append(movies, movie)
	}
	return movies, nil
}

func (r MovieRepository) FindSimilar(embedding string) ([]domain.Movie, error) {
	// TODO

	movies := []domain.Movie{}
	return movies, nil
}
