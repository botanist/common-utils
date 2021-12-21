package serviceutils

import (
	"fmt"
	"net/url"
)

func buildUrl(svc string, path string, query *url.Values) string {
	addr, ok := addrLookup[svc]
	if !ok {
		addrLookupMutex.Lock()
		addr = fmt.Sprintf("http://%s:8080", svc)
		addrLookup[svc] = addr
		addrLookupMutex.Unlock()
	}

	if query != nil {
		return fmt.Sprintf("%s%s?%s", addr, path, query.Encode())
	} else {
		return fmt.Sprintf("%s%s", addr, path)
	}
}
