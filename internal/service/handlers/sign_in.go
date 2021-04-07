package handlers

import (
	"github.com/one-click-platform/deployer/internal/service/auth"
	"github.com/one-click-platform/deployer/internal/service/requests"
	"github.com/one-click-platform/deployer/internal/service/responses"
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

	result, err := AccountsQ(r).FilterByEmail(request.Data.Attributes.Email).Get()
	if err != nil || result == nil {
		Log(r).WithError(err).Info("login not found")
		ape.RenderErr(w, problems.Forbidden())
		return
	}

	if !auth.PasswordsEqual(result.Password, request.Data.Attributes.Password) {
		Log(r).WithError(err).Info("wrong password")
		ape.RenderErr(w, problems.Forbidden())
		return
	}

	response := responses.NewSignInResponse(result)
	ape.Render(w, response)
}
