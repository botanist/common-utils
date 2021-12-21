package serviceutils

import (
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

var addrLookup map[string]string
var addrLookupMutex sync.RWMutex
var client *http.Client

const EnvSvcAddrSuffix = "_ADDR"

func init() {
	// enumerate env
	addrLookup = getEnvServices()
	client = &http.Client{Timeout: time.Second * 30}
}

func getEnvServices() map[string]string {
	m := make(map[string]string)
	for _, v := range os.Environ() {
		parts := strings.SplitN(v, "=", 2)
		if strings.HasSuffix(parts[0], EnvSvcAddrSuffix) {

			svc := strings.TrimSuffix(parts[0], EnvSvcAddrSuffix)
			svc = strings.ReplaceAll(svc, "_", "-")
			m[strings.ToLower(svc)] = parts[1]
		}
	}

	return m
}
