package yarnsite

import (
	"encoding/json"
	"fetbari/common"
	"fmt"
)

const (
	yarnSiteUrlTmpl = "http://%s:8080/api/v1/clusters/%s/configurations?type=yarn-site"
)

func YarnResourceManager(ambari, username, password, clustername string) (string, error) {
	url := fmt.Sprintf(yarnSiteUrlTmpl, ambari, clustername)

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
	yrm, err2 := getYarnResourceManager(username, password, targetHref)
	if err2 != nil {
		return "", err2
	}
	return yrm, nil
}

func getYarnResourceManager(username, password, url string) (string, error) {
	body, err := common.HttpGet(url, username, password)
	if err != nil {
		return "", err
	}
	var res getYarnResourceManagerResponse
	_ = json.Unmarshal(body, &res)
	return res.Items[0].Properties.YarnResourceManagerAddress, nil
}

type getYarnResourceManagerResponse struct {
	Items []item `json:"items"`
}
type item struct {
	Properties property `json:"properties"`
}
type property struct {
	YarnResourceManagerAddress string `json:"yarn.resourcemanager.webapp.address"`
}
