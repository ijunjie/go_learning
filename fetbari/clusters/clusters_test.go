package clusters

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClusterName(t *testing.T) {
	name, err := ClusterName("10.69.71.33", "admin", "admin")
	t.Logf("cluster name: %s", name)
	assert.Nil(t, err)
}
