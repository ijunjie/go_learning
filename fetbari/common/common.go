package common

import (
	"encoding/base64"
	"errors"
	"io/ioutil"
	"net/http"
)

func HttpGet(url, username, password string) ([]byte, error) {
	req, err1 := http.NewRequest("GET", url, nil)
	if err1 != nil {
		return nil, err1
	}
	if username != "" && password != "" {
		req.SetBasicAuth(username, password)
	}

	cli := &http.Client{}
	resp, err2 := cli.Do(req)
	if err2 != nil {
		return nil, err2
	}

	defer resp.Body.Close()

	body, err3 := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		return nil, errors.New("Status code is " + string(resp.StatusCode))
	}
	return body, err3
}

func Base64Encode(username, password string) string {
	input := []byte(username + ":" + password)
	encodeString := base64.StdEncoding.EncodeToString(input)
	return encodeString
}

func OnlineOrOffline(env string) int {
	if env == "online" {
		return 0
	}
	return 1
}

type SiteConfsResponse struct {
	Items []item `json:"items"`
}

type item struct {
	Href    string `json:"href"`
	Version int    `json:"version"`
}
