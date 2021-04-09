package handlers

import (
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"net/http"

	"github.com/go-chi/chi"
)

func GetEnv(w http.ResponseWriter, r *http.Request) {
	// TODO add validation
	name := chi.URLParam(r, "name")

	//accountID, err := helpers.ParsePayload(JWTPayload(r))
	//if err != nil {
	//	Log(r).WithError(err).Info("wrong request")
	//	ape.RenderErr(w, problems.BadRequest(err)...)
	//	return
	//}

	//result, err := EnvsQ(r).FilterByAccountID(accountID).FilterByName(name).Get()
	env, ok := Storage(r)[name]
	if !ok {
		ape.RenderErr(w, problems.NotFound())
		return
	}

	ape.Render(w, env)
}
