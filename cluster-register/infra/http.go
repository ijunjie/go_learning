package infra

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// HttpGet http get with auth
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
		return nil, errors.New("Status code is " + fmt.Sprint(resp.StatusCode))
	}
	return body, err3
}
