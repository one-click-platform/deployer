package handlers

import (
	"github.com/one-click-platform/deployer/internal/data"
	"github.com/one-click-platform/deployer/internal/service/helpers"
	"github.com/one-click-platform/deployer/internal/service/responses"
	"net/http"

	"github.com/one-click-platform/deployer/internal/service/requests"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func CreateNode(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewCreateNodeRequest(r)
	if err != nil {
		Log(r).WithError(err).Info("wrong request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	accountID, err := helpers.ParsePayload(JWTPayload(r))
	if err != nil {
		Log(r).WithError(err).Info("wrong request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	Tasks(r) <- request.Data.Attributes.Name

	result, err := EnvsQ(r).Insert(data.Env{
		Name:      request.Data.Attributes.Name,
		AccountID: accountID,
	})
	if err != nil {
		Log(r).WithError(err).Info("can't insert info")
		ape.RenderErr(w, problems.Conflict())
		return
	}

	response := responses.NewCreateNodeResponse(result, data.Account{ID: accountID})
	ape.Render(w, response)
}
