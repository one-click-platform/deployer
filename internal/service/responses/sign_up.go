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
				AccessToken: "",
				Email:       account.Email,
				Expiry:      0,
				Verified:    false,
			},
		},
	}

	return result
}
