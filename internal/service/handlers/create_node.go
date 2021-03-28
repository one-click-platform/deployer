package handlers

import (
	"net/http"

	"github.com/one-click-platform/deployer/resources"

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

	storage := Storage(r)

	if _, ok := storage[request.Name]; ok {
		ape.RenderErr(w, problems.Conflict())
		return
	}

	storage[request.Name] = resources.EnvConfig{
		Status: "processing",
	}

	Tasks(r) <- request.Name

	w.WriteHeader(http.StatusNoContent)
}
