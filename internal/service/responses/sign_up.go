package responses

import (
	"github.com/one-click-platform/deployer/internal/data"
	"github.com/one-click-platform/deployer/resources"
)

func NewSignUpResponse(account data.Account) resources.AccountObjectResponse {
	result := resources.AccountObjectResponse{
		Data: resources.AccountObject{
			Key: resources.Key{},
			Attributes: resources.AccountObjectAttributes{
				Email:    account.Email,
				Verified: false,
			},
		},
	}

	return result
}
