package rpc

import (
	"encoding/json"
	"log"
	"net/http"
)

func SendJSON(res http.ResponseWriter, statusCode int, out interface{}) {
	res.WriteHeader(statusCode)
	err := json.NewEncoder(res).Encode(out)
	if err != nil {
		log.Println(err)
	}
}
