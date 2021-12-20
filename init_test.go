package serviceutils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindAddrFromEnvironment(t *testing.T) {
	os.Clearenv()
	m := getEnvServices()
	assert.Empty(t, m)

	os.Setenv("AUTH_ADDR", "http://localhost:3000")
	os.Setenv("ACCOUNT_ADDR", "http://localhost:3001")
	os.Setenv("SITE_ADDR", "http://localhost:3002")

	m = getEnvServices()
	assert.NotEmpty(t, m)
}
