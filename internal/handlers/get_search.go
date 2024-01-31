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

	if err != nil {
		/*TODO*/
	}

	pages.Search(movies).Render(context.Background(), w)
}
