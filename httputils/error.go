package httputils

import (
	"encoding/json"
	"log"
	"net/http"
)

type RemoteError struct {
	Ok         bool   `json:"ok"`
	Msg        string `json:"error"`
	StatusCode int
}

type NoError struct {
	OK  bool   `json:"ok"`
	Msg string `json:"msg"`
}

func (r RemoteError) Error() string {
	return r.Msg
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
