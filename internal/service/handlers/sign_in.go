package handlers

import (
	"errors"
	"github.com/one-click-platform/deployer/internal/data"
	"github.com/one-click-platform/deployer/internal/service/requests"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"net/http"
)

func SignIn(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewSignInRequest(r)

	if err != nil {
		Log(r).WithError(err).Info("wrong request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	result, err := AccountsQ(r).FilterByEmail(request.Data.Email).Get()
	// TODO change to error schema
	if err != nil || !data.ComparePassword(result.Password, request.Data.Password) {
		Log(r).WithError(err).Info("login or password not found")
		ape.RenderErr(w, problems.NotAllowed(errors.New("no match for login/password pair")))
		return
	}

	ape.Render(w, result)
}
