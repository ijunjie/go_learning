package main

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	argCount := len(os.Args[1:])
	if argCount != 4 {
		fmt.Println("Usage: fetbari {ambari server ip} {username} {password} {online/offline}")
		fmt.Println("Example: fetbari 10.10.10.10 admin admin online")
		return
	}
	ambari := os.Args[1]
	username := os.Args[2]
	password := os.Args[3]
	env := os.Args[4]
	if env != "online" && env != "offline" {
		fmt.Println("The env arg should be online/offline")
		return
	}

	clusterName, _ := clusterName(ambari, username, password)
	yarnResourceManager, _ := yarnResourceManager(ambari, username, password, clusterName)
	nameNodeHost, _ := nameNodeHost(ambari, username, password, clusterName)
	metricsURL := fmt.Sprintf("http://%s/ws/v1/cluster/metrics", yarnResourceManager)
	vcores, _ := vcores(metricsURL)

	fmt.Printf("%21s: %-30s\n", "Cluster Name", clusterName)
	host := fmt.Sprintf("http://%s:8080", ambari)
	fmt.Printf("%21s: %-30s\n", "Host", host)
	basicKey := base64Encode(username, password)
	fmt.Printf("%21s: %-30s\n", "Authorization", basicKey)
	fmt.Printf("%21s: %-30s\n", "Yarn Resource Manager", yarnResourceManager)
	fmt.Printf("%21s: %-30s\n", "Name Node Host", nameNodeHost)
	fmt.Printf("%21s: %-30d\n", "Total Vcores", vcores)
	envNumber := onlineOrOffline(env)
	fmt.Printf("%21s: %-30d\n", "Env", envNumber)

	format := `INSERT INTO cluster_config (cluster_name,host,root_cu_num,basic_key,rm_host,nm_host,is_default,timestamp,file_type,cluster_type,cluster_kind,hadoop_master_ip,ingress) VALUES ('%s','%s',%d,'%s','%s','%s',1,now(),'core-site,hdfs-site,hive-site,yarn-site,spark2-defaults',%d,0,'%s','{"kdm":"http://some-addr"}');`
	sql := fmt.Sprintf(format, clusterName, host, vcores, basicKey, yarnResourceManager, nameNodeHost, envNumber, ambari)
	fmt.Printf("%21s: %-30s\n", "Insert SQL Sample", sql)
}

func onlineOrOffline(env string) int {
	if env == "online" {
		return 0
	}
	return 1
}

func base64Encode(username, password string) string {
	input := []byte(username + ":" + password)
	encodeString := base64.StdEncoding.EncodeToString(input)
	return encodeString
}

func httpGet(url, username, password string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	if username != "" && password != "" {
		req.SetBasicAuth(username, password)
	}

	cli := &http.Client{}
	resp, err := cli.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New("Status code is " + string(resp.StatusCode))
	}
	return body, nil
}

func clusterName(ambari, username, password string) (string, error) {
	url := fmt.Sprintf("http://%s:8080/api/v1/clusters", ambari)
	body, err := httpGet(url, username, password)
	if err != nil {
		return "", err
	}

	var clustersResponse clusterResponse
	_ = json.Unmarshal(body, &clustersResponse)
	return clustersResponse.Items[0].Clusters.ClusterName, nil
}

type cluster struct {
	ClusterName string `json:"cluster_name"`
}
type item struct {
	Clusters cluster `json:"Clusters"`
}

type clusterResponse struct {
	Items []item `json:"items"`
}

func yarnResourceManager(ambari, username, password, clustername string) (string, error) {
	url := fmt.Sprintf("http://%s:8080/api/v1/clusters/%s/configurations?type=yarn-site",
		ambari, clustername)

	body, err := httpGet(url, username, password)
	if err != nil {
		return "", err
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
		}
	}
	yrm, err := getYarnResourceManager(username, password, targetHref)
	if err != nil {
		return "getYarnResourceManager err", err
	}
	return yrm, nil
}

type versionItem struct {
	Href    string `json:"href"`
	Version int    `json:"version"`
}

type siteConfsResponse struct {
	Items []versionItem `json:"items"`
}

func getYarnResourceManager(username, password, url string) (string, error) {
	body, err := httpGet(url, username, password)
	if err != nil {
		return "", err
	}
	var res getYarnResourceManagerResponse
	_ = json.Unmarshal(body, &res)
	return res.Items[0].Properties.YarnResourceManagerAddress, nil
}

type property struct {
	YarnResourceManagerAddress string `json:"yarn.resourcemanager.webapp.address"`
}
type item3 struct {
	Properties property `json:"properties"`
}
type getYarnResourceManagerResponse struct {
	Items []item3 `json:"items"`
}

func nameNodeHost(ambari, username, password, clustername string) (string, error) {
	url := fmt.Sprintf("http://%s:8080/api/v1/clusters/%s/configurations?type=core-site",
		ambari, clustername)
	body, err := httpGet(url, username, password)
	if err != nil {
		return "", err
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
		}
	}
	nameNode, err := getDefaultFS(username, password, targetHref)
	if err != nil {
		return "getDefaultFS err", err
	}
	return nameNode, nil
}

func getDefaultFS(username, password, url string) (string, error) {
	body, err := httpGet(url, username, password)
	if err != nil {
		return "", err
	}
	var res getCoreSiteResponse
	_ = json.Unmarshal(body, &res)
	return res.Items[0].Properties.DefaultFS, nil
}

type properties struct {
	DefaultFS string `json:"fs.defaultFS"`
}
type item4 struct {
	Properties properties `json:"properties"`
}
type getCoreSiteResponse struct {
	Items []item4 `json:"items"`
}

func vcores(url string) (int, error) {
	body, err := httpGet(url, "", "")
	if err != nil {
		return 0, err
	}
	var res metricsResponse
	_ = json.Unmarshal(body, &res)
	return res.ClusterMetrics.TotalVirtualCores, nil
}

type clusterMetrics struct {
	TotalVirtualCores int `json:"totalVirtualCores"`
}
type metricsResponse struct {
	ClusterMetrics clusterMetrics `json:"clusterMetrics"`
}
