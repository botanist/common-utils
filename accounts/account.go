package accounts

import (
	"fmt"
	"net/http"

	"github.com/botanist/common-utils/rpc"
)

type Account struct {
	Ok                    bool     `json:"ok"`
	Error                 string   `json:"error"`
	ID                    string   `json:"id"`
	Sites                 []string `json:"sites"`
	Email                 string   `json:"email",omitempty`
	Verified              bool     `json:"verified"`
	VerificationChallenge string   `json:"verification_challenge",omitempty`
	Auths                 []string `json:"auths`
}

func GetAccount(jwt string, id string) (*Account, int, error) {
	var account Account
	rs, err := rpc.GetJSON("api.account", fmt.Sprintf("/v1/account/%s", id), nil, jwt, nil, &account)
	if rs != http.StatusOK {
		return nil, rs, err
	}

	return &account, rs, nil
}
