package handlers

import (
	"github.com/one-click-platform/deployer/internal/data"
	"github.com/one-click-platform/deployer/internal/service/helpers"
	"github.com/one-click-platform/deployer/internal/service/responses"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"net/http"

	"github.com/go-chi/chi"
)

func GetEnv(w http.ResponseWriter, r *http.Request) {
	// TODO add validation
	name := chi.URLParam(r, "name")

	accountID, err := helpers.ParsePayload(JWTPayload(r))
	if err != nil {
		Log(r).WithError(err).Info("wrong request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	environment, err := EnvsQ(r).FilterByAccountID(accountID).FilterByName(name).Get()
	if err != nil {
		ape.RenderErr(w, problems.NotFound())
		return
	}

	response := responses.NewGetEnvResponse(*environment, data.Account{ID: accountID})
	ape.Render(w, response)
}
