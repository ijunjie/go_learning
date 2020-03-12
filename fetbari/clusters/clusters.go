package clusters

import (
	"encoding/json"
	"fetbari/common"
	"fmt"
)

const (
	clusterUrlTmpl = "http://%s:8080/api/v1/clusters"
)

func ClusterName(ambari, username, password string) (string, error) {
	url := fmt.Sprintf(clusterUrlTmpl, ambari)
	body, err := common.HttpGet(url, username, password)
	if err != nil {
		return "", err
	}

	var clustersResponse clusterResponse
	_ = json.Unmarshal(body, &clustersResponse)
	return clustersResponse.Items[0].Clusters.ClusterName, nil
}

type clusterResponse struct {
	Items []item `json:"items"`
}
type item struct {
	Clusters cluster `json:"Clusters"`
}
type cluster struct {
	ClusterName string `json:"cluster_name"`
}
