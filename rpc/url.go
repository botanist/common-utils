package rpc

import (
	"fmt"
	"net/url"
)

func buildUrl(svc string, path string, query *url.Values) string {
	addr := fmt.Sprintf("http://%s", svc)
	if query != nil {
		return fmt.Sprintf("%s%s?%s", addr, path, query.Encode())
	}

	return fmt.Sprintf("%s%s", addr, path)
}
