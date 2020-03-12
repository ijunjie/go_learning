package yarnsite

import "testing"

func TestYarnResourceManager(t *testing.T) {
	yrm, err := YarnResourceManager("10.69.71.33", "admin", "admin", "online")
	if err != nil {
		t.Error(err)
	}
	t.Logf("yrm: %s", yrm)
}
