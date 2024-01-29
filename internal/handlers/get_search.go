package handlers

import (
	"context"
	"net/http"
	"net/url"

	"github.com/brunocapri/movie-engine/internal/templates/pages"
)

func (h Handler) GetSearch(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query().Get("q")
	if param == "" { /*TODO*/
	}
	query, err := url.QueryUnescape(param)

	if err != nil { /*TODO*/
	}

	res, err := h.app.OpenAiClient.GenerateEmbeddings(query)

	if err != nil { /*TODO*/
	}

	embedding := h.app.OpenAiClient.EmbeddingToString(res.Data[0].Embedding)

	movies, err := h.movies.FindByEmbedding(embedding)

	pages.Search(movies).Render(context.Background(), w)
}
