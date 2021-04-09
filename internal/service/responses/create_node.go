package responses

import (
	"github.com/one-click-platform/deployer/internal/data"
	"github.com/one-click-platform/deployer/resources"
)

func NewCreateNodeResponse(env data.Env, account data.Account) resources.EnvironmentResponse {
	return resources.EnvironmentResponse{
		Data: resources.Environment{
			Key:        resources.NewKeyInt64(env.ID, resources.ENVIRONMENT),
			Attributes: resources.EnvironmentAttributes{Name: env.Name},
			Relationships: &resources.EnvironmentRelationships{
				Data: resources.RelationCollection{
					Data: []resources.Key{
						resources.NewKeyInt64(account.ID, resources.ACCOUNT),
					},
				},
			},
		},
		Included: resources.Included{},
	}
}
