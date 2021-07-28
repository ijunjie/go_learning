package kde

import (
	"cluster-register/infra"
	"encoding/base64"
	"encoding/json"
	"fmt"
)

const (
	clusterUrlTmpl  = "http://%s/api/v1/clusters"
	yarnSiteUrlTmpl = "http://%s/api/v1/clusters/%s/configurations?type=yarn-site"
	coreSiteUrlTmpl = "http://%s/api/v1/clusters/%s/configurations?type=core-site"
	metricsUrlTmpl  = "http://%s/ws/v1/cluster/metrics"
)

type KdeInfoRequest struct {
	KdeHost  string
	KdePort  int
	Username string
	Password string
	KdeType  string
}

type KdeInfoResult struct {
	ClusterName         string
	Host                string
	BasicKey            string
	YarnResourceManager string
	NameNodeHost        string
	Vcores              int
	Mem                 int
	MemGB               int
	Type                string
	Env                 int
	HadoopMasterIp      string
}

func (info *KdeInfoResult) ToClusterConfigInsert() *infra.ClusterConfigInsert {
	return &infra.ClusterConfigInsert{
		ClusterName:    info.ClusterName,
		Host:           info.Host,
		RootCuNum:      info.Vcores,
		BasicKey:       info.BasicKey,
		RmHost:         info.YarnResourceManager,
		NmHost:         info.NameNodeHost,
		ClusterType:    info.Env,
		HadoopMasterIp: info.HadoopMasterIp,
	}
}

func KdeInfo(request *KdeInfoRequest) (*KdeInfoResult, error) {
	// kdeHost string, kdePort int, username, password, kdeType string
	kdeHost, kdePort, username, password, kdeType := request.KdeHost, request.KdePort, request.Username, request.Password, request.KdeType
	ambari := fmt.Sprintf("%s:%d", kdeHost, kdePort)
	clusterName, err1 := clusterName(ambari, username, password)
	if err1 != nil {
		return nil, err1
	}

	yarnResourceManager, err2 := yarnResourceManager(ambari, username, password, clusterName)
	if err2 != nil {
		return nil, err2
	}

	nameNodeHost, err3 := nameNodeHost(ambari, username, password, clusterName)
	if err3 != nil {
		return nil, err3
	}
	host := fmt.Sprintf("http://%s", ambari)
	basicKey := base64Encode(username, password)

	vcores, mem, err4 := vcoresAndMem(yarnResourceManager)
	if err4 != nil {
		return nil, err4
	}
	memGB := mem / 1024
	envNumber := onlineOrOffline(kdeType)

	return &KdeInfoResult{
		ClusterName:         clusterName,
		Host:                host,
		BasicKey:            basicKey,
		YarnResourceManager: yarnResourceManager,
		NameNodeHost:        nameNodeHost,
		Vcores:              vcores,
		Mem:                 mem,
		MemGB:               memGB,
		Type:                kdeType,
		Env:                 envNumber,
		HadoopMasterIp:      kdeHost,
	}, nil
}

func base64Encode(username, password string) string {
	input := []byte(username + ":" + password)
	encodeString := base64.StdEncoding.EncodeToString(input)
	return encodeString
}

func onlineOrOffline(env string) int {
	if env == "online" {
		return 0
	}
	return 1
}

// parse cluster name from ambari-api
func clusterName(ambari, username, password string) (string, error) {
	url := fmt.Sprintf(clusterUrlTmpl, ambari)
	body, err := infra.HttpGet(url, username, password)
	if err != nil {
		return "", err
	}

	var clustersResponse clusterResponse
	_ = json.Unmarshal(body, &clustersResponse)
	return clustersResponse.Items[0].Clusters.ClusterName, nil
}

type clusterResponse struct {
	Items []item1 `json:"items"`
}
type item1 struct {
	Clusters cluster `json:"Clusters"`
}
type cluster struct {
	ClusterName string `json:"cluster_name"`
}

// --- fetch yarn server
func yarnResourceManager(ambari, username, password, clustername string) (string, error) {
	url := fmt.Sprintf(yarnSiteUrlTmpl, ambari, clustername)

	body, err1 := infra.HttpGet(url, username, password)
	if err1 != nil {
		return "", err1
	}

	var siteConfsResponse siteConfsResponse
	_ = json.Unmarshal(body, &siteConfsResponse)
	targetHref := ""
	maxVersion := 1
	for i := 0; i < len(siteConfsResponse.Items); i++ {
		item := siteConfsResponse.Items[i]
		if maxVersion < item.Version {
			maxVersion = item.Version
			targetHref = item.Href
		} else {
			targetHref = item.Href
		}
	}
	yrm, err2 := fetchYrnResourceManager(username, password, targetHref)
	if err2 != nil {
		return "", err2
	}
	return yrm, nil
}

func fetchYrnResourceManager(username, password, url string) (string, error) {
	body, err := infra.HttpGet(url, username, password)
	if err != nil {
		return "", err
	}
	var res yarnResourceManagerResponse
	_ = json.Unmarshal(body, &res)
	return res.Items[0].Properties.YarnResourceManagerAddress, nil
}

type yarnResourceManagerResponse struct {
	Items []item2 `json:"items"`
}
type item2 struct {
	Properties property `json:"properties"`
}
type property struct {
	YarnResourceManagerAddress string `json:"yarn.resourcemanager.webapp.address"`
}

// --- namenode
func nameNodeHost(ambari, username, password, clustername string) (string, error) {
	url := fmt.Sprintf(coreSiteUrlTmpl, ambari, clustername)
	body, err1 := infra.HttpGet(url, username, password)
	if err1 != nil {
		return "", err1
	}
	var siteConfsResponse siteConfsResponse
	_ = json.Unmarshal(body, &siteConfsResponse)
	targetHref := ""
	maxVersion := 1
	for i := 0; i < len(siteConfsResponse.Items); i++ {
		item := siteConfsResponse.Items[i]
		if maxVersion < item.Version {
			maxVersion = item.Version
			targetHref = item.Href
		} else {
			targetHref = item.Href
		}
	}
	nameNode, err2 := fetchDefaultFS(username, password, targetHref)
	if err2 != nil {
		return "", err2
	}
	return nameNode, nil
}

func fetchDefaultFS(username, password, url string) (string, error) {
	body, err := infra.HttpGet(url, username, password)
	if err != nil {
		return "", err
	}
	var res coreSiteResponse
	_ = json.Unmarshal(body, &res)
	return res.Items[0].Properties.DefaultFS, nil
}

type coreSiteResponse struct {
	Items []item3 `json:"items"`
}
type item3 struct {
	Properties properties `json:"properties"`
}
type properties struct {
	DefaultFS string `json:"fs.defaultFS"`
}

type siteConfsResponse struct {
	Items []item0 `json:"items"`
}

type item0 struct {
	Href    string `json:"href"`
	Version int    `json:"version"`
}

// --- vcores and mem
func vcoresAndMem(yrm string) (int, int, error) {
	url := fmt.Sprintf(metricsUrlTmpl, yrm)
	body, err := infra.HttpGet(url, "", "")
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
