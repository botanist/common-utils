package entities

import (
	"fmt"
	"net/http"

	"github.com/botanist/common-utils/rpc"
)

type Entity struct {
	Ok          bool   `json:"ok"`
	Error       string `json:"error"`
	ID          string `json:"id"`
	Site        string `json:"site"`
	Type        string `json:"type"`
	Parent      string `json:"parent"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Location    string `json:"location"`
	Preferences string `json:"preferences"`
}

func GetEntity(jwt string, site string, id string) (*Entity, int, error) {
	var entity Entity
	rs, err := rpc.GetJSON("api.entity", fmt.Sprintf("/v1/entity/%s/%s", site, id), nil, jwt, nil, &entity)
	if rs != http.StatusOK {
		return nil, rs, err
	}

	return &entity, rs, nil
}
