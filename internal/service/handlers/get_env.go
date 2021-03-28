package handlers

import (
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"

	"github.com/go-chi/chi"
)

func GetEnv(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")

	env, ok := Storage(r)[name]
	if !ok {
		ape.RenderErr(w, problems.NotFound())
		return
	}

	ape.Render(w, env)
}
