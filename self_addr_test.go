package serviceutils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeServiceEnv(t *testing.T) {
	os.Setenv("TEST_SERVICE_ADDR", "http://localhost:3000")

	ok, host, port := DecodeServiceEnv("test_service")
	assert.True(t, ok)
	assert.Equal(t, "localhost", host)
	assert.Equal(t, 3000, port)

}
