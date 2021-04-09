package responses

import (
	"github.com/one-click-platform/deployer/internal/data"
	"github.com/one-click-platform/deployer/resources"
)

func NewSignUpResponse(account data.Account, token *string) resources.AccountObjectResponse {
	result := resources.AccountObjectResponse{
		Data: resources.AccountObject{
			Key: resources.NewKeyInt64(account.ID, resources.ACCOUNT),
			Attributes: resources.AccountObjectAttributes{
				Email:       account.Email,
				AccessToken: token,
				Verified:    false,
			},
		},
	}

	return result
}
