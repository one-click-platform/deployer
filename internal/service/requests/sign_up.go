package requests

import (
	"encoding/json"
	"github.com/one-click-platform/deployer/resources"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net/http"
)

type SignUpRequest struct {
	Data resources.Account
}

func NewSignUpRequest(r *http.Request) (SignUpRequest, error) {
	var request SignUpRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return request, errors.Wrap(err, "failed to unmarshal")
	}

	// TODO input validation
	return request, nil
}
