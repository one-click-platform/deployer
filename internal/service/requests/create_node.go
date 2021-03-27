package requests

import (
	"encoding/json"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net/http"
)

type CreateNodeRequest struct {
	Name string `json:"name"`
}

func NewCreateNodeRequest(r *http.Request) (CreateNodeRequest, error) {
	var request CreateNodeRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return request, errors.Wrap(err, "failed to unmarshal")
	}

	return request, nil
}
