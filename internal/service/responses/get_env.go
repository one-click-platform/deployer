package responses

import (
	"github.com/one-click-platform/deployer/internal/data"
	"github.com/one-click-platform/deployer/resources"
)

func NewGetEnvResponse(env data.Env, account data.Account) resources.EnvironmentResponse {
	return resources.EnvironmentResponse{
		Data: resources.Environment{
			Key:           resources.Key{},
			Attributes:    resources.EnvironmentAttributes{},
			Relationships: &resources.EnvironmentRelationships{},
		},
		Included: resources.Included{},
	}
}
