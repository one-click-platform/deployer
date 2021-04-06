package handlers

import (
	"github.com/one-click-platform/deployer/internal/data"
	"github.com/one-click-platform/deployer/internal/service/auth"
	"github.com/one-click-platform/deployer/internal/service/requests"
	"github.com/one-click-platform/deployer/internal/service/responses"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"net/http"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewSignUpRequest(r)
	if err != nil {
		Log(r).WithError(err).Info("wrong request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	hashedPassword, err := auth.HashPassword(request.Data.Attributes.Password)
	if err != nil {
		Log(r).WithError(err).Info("wrong password")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	result, err := AccountsQ(r).Insert(data.Account{
		Email:    request.Data.Attributes.Email,
		Password: hashedPassword,
	})
	if err != nil {
		Log(r).WithError(err).Info("Failed to create account")
		ape.RenderErr(w, problems.Conflict())
		return
	}

	response := responses.NewSignUpResponse(result)
	ape.Render(w, response)
}
