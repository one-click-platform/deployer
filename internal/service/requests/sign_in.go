package requests

import (
	"encoding/json"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/one-click-platform/deployer/resources"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net/http"
)

type SignInRequest struct {
	Data resources.Account
}

func NewSignInRequest(r *http.Request) (SignInRequest, error) {
	var request SignInRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return request, errors.Wrap(err, "failed to unmarshal")
	}

	// TODO add validation rules
	return request, nil
}

func (r *SignInRequest) validate() error {
	return mergeErrors(validation.Errors{}).Filter()
}

func mergeErrors(validationErrors ...validation.Errors) validation.Errors {
	result := make(validation.Errors)
	for _, errs := range validationErrors {
		for key, err := range errs {
			result[key] = err
		}
	}
	return result
}
