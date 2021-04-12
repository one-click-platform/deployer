package responses

import (
	"github.com/one-click-platform/deployer/internal/data"
	"github.com/one-click-platform/deployer/resources"
)

func NewSignInResponse(account *data.Account, token string) resources.AccountObjectResponse {

	result := resources.AccountObjectResponse{
		Data: resources.AccountObject{
			Key: resources.Key{},
			Attributes: resources.AccountObjectAttributes{
				AccessToken: token,
				Email:       account.Email,
				Expiry:      0,
				Verified:    false,
			},
		},
	}

	return result
}
