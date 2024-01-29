package handlers

import (
	"context"
	"net/http"

	"github.com/brunocapri/movie-engine/internal/templates/pages"
)

func (h Handler) GetHome(w http.ResponseWriter, r *http.Request) {
	pages.Home().Render(context.Background(), w)
}
