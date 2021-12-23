package serviceutils

import (
	"encoding/json"
	"log"
	"net/http"
)

type RemoteError struct {
	StatusCode int
	OK         bool   `json:"ok"`
	ErrorMsg   string `json:"error"`
}

type NoError struct {
	OK bool `json:"ok"`
}

func (r RemoteError) Error() string {
	return r.ErrorMsg
}

func SendError(res http.ResponseWriter, statusCode int, err error) {
	e := struct {
		OK    bool   `json:"ok"`
		Error string `json:"error"`
	}{false, err.Error()}

	res.WriteHeader(statusCode)
	err = json.NewEncoder(res).Encode(e)
	if err != nil {
		log.Println(err)
	}
}
