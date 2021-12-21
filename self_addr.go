package serviceutils

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func DecodeServiceEnv(s string) (bool, string, int) {
	s = strings.ReplaceAll(s, "-", "_")
	k := strings.ToUpper(fmt.Sprintf("%s_ADDR", s))
	v := os.Getenv(k)

	u, err := url.Parse(v)
	if err != nil {
		log.Println(err)
		return false, "", 0
	}

	port := 0
	if p, err := strconv.Atoi(u.Port()); err == nil {
		port = p
	}

	return true, u.Hostname(), port
}
