package repository

import (
	"database/sql"
	"log"

	"github.com/brunocapri/movie-engine/internal/domain"
	"github.com/google/uuid"
)

type MovieRepository struct {
	db *sql.DB
}

func NewMovieRepository(db *sql.DB) *MovieRepository {
	return &MovieRepository{db: db}
}

func (r MovieRepository) FindByEmbedding(embedding string) ([]domain.Movie, error) {
	const query = "SELECT id, title, year, plot, runtime, director, poster, ratings FROM omdb_movies ORDER BY embedding <-> $1 LIMIT 5"
	rows, err := r.db.Query(query, embedding)
	if err != nil {
		log.Printf("Query error: %v", err)
		return nil, err
	}
	defer rows.Close()

	movies := []domain.Movie{}
	for rows.Next() {
		movie := domain.Movie{}
		err := rows.Scan(&movie.Id, &movie.Title, &movie.Year, &movie.Plot, &movie.Runtime, &movie.Director, &movie.Poster, &movie.Ratings)
		if err != nil {
			log.Printf("Scan error: %v", err)
			continue
		}
		movies = append(movies, movie)
	}
	return movies, nil
}

func (r MovieRepository) FindSimilar(movie domain.Movie) ([]domain.Movie, error) {
	const query = "SELECT id, title, year, plot, runtime, director, poster, ratings FROM omdb_movies WHERE id != $1 ORDER BY embedding <-> $2 LIMIT 5"
	rows, err := r.db.Query(query, movie.Id, movie.Embedding)
	if err != nil {
		log.Printf("Query error: %v", err)
		return nil, err
	}
	defer rows.Close()

	movies := []domain.Movie{}
	for rows.Next() {
		movie := domain.Movie{}
		err := rows.Scan(&movie.Id, &movie.Title, &movie.Year, &movie.Plot, &movie.Runtime, &movie.Director, &movie.Poster, &movie.Ratings)
		if err != nil {
			log.Printf("Scan error: %v", err)
			continue
		}
		movies = append(movies, movie)
	}
	return movies, nil
}

func (r MovieRepository) FindByTitle(title string) (domain.Movie, error) {
	movie := domain.Movie{}
	const query = "SELECT id, title, year, plot, runtime, director, poster, ratings, embedding FROM omdb_movies WHERE UPPER(title) LIKE UPPER($1)"
	titleSearch := "%" + title + "%"
	err := r.db.QueryRow(query, titleSearch).Scan(&movie.Id, &movie.Title, &movie.Year, &movie.Plot, &movie.Runtime, &movie.Director, &movie.Poster, &movie.Ratings, &movie.Embedding)

	return movie, err
}

func (r MovieRepository) FindById(id uuid.UUID) (domain.Movie, error) {
	movie := domain.Movie{}
	const query = "SELECT id, title, year, plot, runtime, director, poster, ratings, embedding FROM omdb_movies WHERE id = $1"
	err := r.db.QueryRow(query, id).Scan(&movie.Id, &movie.Title, &movie.Year, &movie.Plot, &movie.Runtime, &movie.Director, &movie.Poster, &movie.Ratings, &movie.Embedding)

	return movie, err
}
