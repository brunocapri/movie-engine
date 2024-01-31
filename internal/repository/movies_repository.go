package repository

import (
	"database/sql"
	"encoding/json"
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
		var ratingsRaw json.RawMessage
		err := rows.Scan(&movie.Id, &movie.Title, &movie.Year, &movie.Plot, &movie.Runtime, &movie.Director, &movie.Poster, &ratingsRaw)
		movie.Ratings = parseRatings(ratingsRaw)
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
		var ratingsRaw json.RawMessage
		err := rows.Scan(&movie.Id, &movie.Title, &movie.Year, &movie.Plot, &movie.Runtime, &movie.Director, &movie.Poster, &ratingsRaw)
		movie.Ratings = parseRatings(ratingsRaw)
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
	var ratingsRaw json.RawMessage
	err := r.db.QueryRow(query, titleSearch).Scan(&movie.Id, &movie.Title, &movie.Year, &movie.Plot, &movie.Runtime, &movie.Director, &movie.Poster, &ratingsRaw, &movie.Embedding)
	movie.Ratings = parseRatings(ratingsRaw)

	return movie, err
}

func (r MovieRepository) FindById(id uuid.UUID) (domain.Movie, error) {
	movie := domain.Movie{}
	var rating json.RawMessage
	const query = "SELECT id, title, year, plot, runtime, director, poster, ratings, embedding FROM omdb_movies WHERE id = $1"
	err := r.db.QueryRow(query, id).Scan(&movie.Id, &movie.Title, &movie.Year, &movie.Plot, &movie.Runtime, &movie.Director, &movie.Poster, &rating, &movie.Embedding)
	movie.Ratings = parseRatings(rating)
	return movie, err
}

func parseRatings(ratingsRaw json.RawMessage) []domain.Rating {
	var ratings []domain.Rating
	json.Unmarshal(ratingsRaw, &ratings)
	return ratings
}
