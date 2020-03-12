package coresite

import (
	"encoding/json"
	"fetbari/common"
	"fmt"
)

const (
	coreSiteUrlTmpl = "http://%s:8080/api/v1/clusters/%s/configurations?type=core-site"
)

func NameNodeHost(ambari, username, password, clustername string) (string, error) {
	url := fmt.Sprintf(coreSiteUrlTmpl, ambari, clustername)
	body, err1 := common.HttpGet(url, username, password)
	if err1 != nil {
		return "", err1
	}
	var siteConfsResponse common.SiteConfsResponse
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
	nameNode, err2 := getDefaultFS(username, password, targetHref)
	if err2 != nil {
		return "", err2
	}
	return nameNode, nil
}

func getDefaultFS(username, password, url string) (string, error) {
	body, err := common.HttpGet(url, username, password)
	if err != nil {
		return "", err
	}
	var res getCoreSiteResponse
	_ = json.Unmarshal(body, &res)
	return res.Items[0].Properties.DefaultFS, nil
}

type getCoreSiteResponse struct {
	Items []item `json:"items"`
}
type item struct {
	Properties properties `json:"properties"`
}
type properties struct {
	DefaultFS string `json:"fs.defaultFS"`
}
