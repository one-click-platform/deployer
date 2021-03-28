package requests

import (
	"encoding/json"
	"net/http"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation"

	"gitlab.com/distributed_lab/logan/v3/errors"
)

type CreateNodeRequest struct {
	Name string `json:"name"`
}

func NewCreateNodeRequest(r *http.Request) (CreateNodeRequest, error) {
	var request CreateNodeRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return request, errors.Wrap(err, "failed to unmarshal")
	}

	request.Name = strings.ReplaceAll(request.Name, " ", "")

	return request, request.validate()
}

func (r *CreateNodeRequest) validate() error {
	return validation.Errors{
		"name": validation.Validate(&r.Name, validation.Required, validation.Length(1, 100)),
	}.Filter()
}
