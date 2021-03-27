package handlers

import (
	"net/http"

	"github.com/one-click-platform/deployer/internal/deploy"
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

	_, err = deploy.DeployEC2(request.Name, Log(r))
	if err != nil {
		Log(r).WithError(err).Error("failed to deploy node")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
