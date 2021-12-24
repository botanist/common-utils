package serviceutils

import (
	"encoding/json"
	"log"
	"net/http"
)

func SendOK(res http.ResponseWriter, statusCode int, msg string) {
	e := struct {
		Ok  bool   `json:"ok"`
		Msg string `json:"msg"`
	}{true, msg}

	res.WriteHeader(statusCode)
	err := json.NewEncoder(res).Encode(e)
	if err != nil {
		log.Println(err)
	}
}
