package kde

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClusterName(t *testing.T) {
	ambari := fmt.Sprintf("%s:%d", "10.69.75.29", 8080)
	name, err := clusterName(ambari, "admin", "admin", 2)
	t.Logf("cluster name: %s", name)
	assert.Nil(t, err)
}
