package requests

import (
	"encoding/json"
	"github.com/one-click-platform/deployer/resources"
	"net/http"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation"

	"gitlab.com/distributed_lab/logan/v3/errors"
)

type CreateNodeRequest struct {
	Data resources.Environment
}

func NewCreateNodeRequest(r *http.Request) (CreateNodeRequest, error) {
	var request CreateNodeRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return request, errors.Wrap(err, "failed to unmarshal")
	}

	request.Data.Attributes.Name = strings.ReplaceAll(request.Data.Attributes.Name, " ", "")

	return request, request.validate()
}

func (r *CreateNodeRequest) validate() error {
	return validation.Errors{
		"data/attributes/name": validation.Validate(&r.Data.Attributes.Name, validation.Required, validation.Length(1, 100)),
	}.Filter()
}
