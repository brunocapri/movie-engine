package handlers

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/brunocapri/movie-engine/internal/domain"
	"github.com/brunocapri/movie-engine/internal/templates/pages"
	"github.com/google/uuid"
)

func (h Handler) GetSearch(w http.ResponseWriter, r *http.Request) {

	from := r.URL.Query().Get("from")

	if from != "" {
		id, _ := uuid.Parse(from)
		movie := domain.Movie{Id: id}
		movies, _ := h.core.GetSimilarMovies(movie, true)
		pages.Search(movies).Render(context.Background(), w)
		return
	}

	param := r.URL.Query().Get("q")
	if param == "" { /*TODO*/
	}

	query, err := url.QueryUnescape(param)

	if err != nil { /*TODO*/
	}

	fmt.Println(query)

	movies, err := h.core.SearchMoviesByEmbedding(query)

	// var movies = []domain.Movie{
	// 	{
	// 		Title:    "The Shawshank Redemption",
	// 		Year:     "1994",
	// 		Runtime:  "142 min",
	// 		Director: "Frank Darabont",
	// 		Poster:   "https://m.media-amazon.com/images/M/MV5BNDE3ODcxYzMtY2YzZC00NmNlLWJiNDMtZDViZWM2MzIxZDYwXkEyXkFqcGdeQXVyNjAwNDUxODI@._V1_SX300.jpg",
	// 	},
	// 	{
	// 		Title:    "Seven Pounds",
	// 		Year:     "2008",
	// 		Runtime:  "123 min",
	// 		Director: "Gabriele Muccino",
	// 		Poster:   "https://m.media-amazon.com/images/M/MV5BMTU0NzY0MTY5OF5BMl5BanBnXkFtZTcwODY3MDEwMg@@._V1_SX300.jpg",
	// 	},
	// 	// Add more movies in the same format
	// }

	if err != nil {
		/*TODO*/
	}

	pages.Search(movies).Render(context.Background(), w)
}
