package handlers

import (
	"github.com/one-click-platform/deployer/resources"
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

	//accountID, err := helpers.ParsePayload(JWTPayload(r))
	//if err != nil {
	//	Log(r).WithError(err).Info("wrong request")
	//	ape.RenderErr(w, problems.BadRequest(err)...)
	//	return
	//}
	////TODO change sequence after making status field and reworking storage
	//_, err = EnvsQ(r).Insert(data.Env{
	//	Name:      request.Data.Attributes.Name,
	//	AccountID: accountID,
	//})

	storage := Storage(r)

	if _, ok := storage[request.Data.Attributes.Name]; ok {
		ape.RenderErr(w, problems.Conflict())
		return
	}

	storage[request.Data.Attributes.Name] = resources.EnvConfig{
		Status: "processing",
	}

	Tasks(r) <- request.Data.Attributes.Name

	w.WriteHeader(http.StatusNoContent)
}
