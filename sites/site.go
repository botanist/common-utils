package accounts

import (
	"fmt"
	"net/http"

	"github.com/botanist/common-utils/rpc"
)

type Site struct {
	Ok    bool   `json:"ok"`
	Error string `json:"error"`
	ID    string `json:"id"`
	Owner string `json:"owner"`
	Shard string `json:"sid"`
}

func GetSite(jwt string, id string) (*Site, int, error) {
	var site Site
	rs, err := rpc.GetJSON("api.site", fmt.Sprintf("/v1/site/%s", id), nil, jwt, nil, &site)
	if rs != http.StatusOK {
		return nil, rs, err
	}

	return &site, rs, nil
}
