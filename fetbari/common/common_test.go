package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOnlineOrOffline(t *testing.T) {
	online := OnlineOrOffline("online")
	t.Logf("env: %d", online)
}

func TestBase64Encode(t *testing.T) {
	encode := Base64Encode("admin", "admin")
	assert.Equal(t, "YWRtaW46YWRtaW4=", encode)
}