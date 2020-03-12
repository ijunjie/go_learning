package metrics

import (
	"encoding/json"
	"fetbari/common"
	"fmt"
)

const (
	metricsUrlTmpl = "http://%s/ws/v1/cluster/metrics"
)

func VcoresAndMem(yrm string) (int, int, error) {
	url := fmt.Sprintf(metricsUrlTmpl, yrm)
	body, err := common.HttpGet(url, "", "")
	if err != nil {
		return 0, 0, err
	}
	var res metricsResponse
	_ = json.Unmarshal(body, &res)
	return res.ClusterMetrics.TotalVirtualCores, res.ClusterMetrics.TotalMB, nil
}

type metricsResponse struct {
	ClusterMetrics clusterMetrics `json:"clusterMetrics"`
}

type clusterMetrics struct {
	TotalVirtualCores int `json:"totalVirtualCores"`
	TotalMB           int `json:"totalMB"`
}
