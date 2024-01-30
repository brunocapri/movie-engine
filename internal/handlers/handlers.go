package handlers

import (
	"github.com/brunocapri/movie-engine/internal/core"
)

type Handler struct {
	core *core.Core
}

func NewHandler(core *core.Core) *Handler {
	return &Handler{
		core: core,
	}
}
