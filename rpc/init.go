package rpc

import (
	"net/http"
	"time"
)

var client *http.Client

func init() {
	// enumerate env
	client = &http.Client{Timeout: time.Second * 30}
}
