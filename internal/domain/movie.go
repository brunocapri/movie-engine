package domain

import (
	"encoding/json"

	"github.com/google/uuid"
)

type Rating struct {
	Source string `json:"Source"`
	Value  string `json:"Value"`
}

type Movie struct {
	Id         uuid.UUID       `json:"id"`
	Title      string          `json:"Title"`
	Year       string          `json:"Year"`
	Rated      string          `json:"Rated"`
	Released   string          `json:"Released"`
	Runtime    string          `json:"Runtime"`
	Genre      string          `json:"Genre"`
	Director   string          `json:"Director"`
	Writer     string          `json:"Writer"`
	Actors     string          `json:"Actors"`
	Plot       string          `json:"Plot"`
	Language   string          `json:"Language"`
	Country    string          `json:"Country"`
	Awards     string          `json:"Awards"`
	Poster     string          `json:"Poster"`
	Ratings    json.RawMessage `json:"Ratings"`
	Metascore  string          `json:"Metascore"`
	ImdbRating string          `json:"ImdbRating"`
	ImdbVotes  string          `json:"imdbVotes"`
	ImdbID     string          `json:"imdbID"`
	Type       string          `json:"Type"`
	Dvd        string          `json:"DVD"`
	BoxOffice  string          `json:"BoxOffice"`
	Production string          `json:"Production"`
	Website    string          `json:"Website"`
	Embedding  string          `json:"embedding"`
}
