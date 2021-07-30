package infra

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// HttpGet http get with auth
func HttpGet(url, username, password string, timeoutSeconds int) ([]byte, error) {
	req, err1 := http.NewRequest("GET", url, nil)
	if err1 != nil {
		return nil, err1
	}
	if username != "" && password != "" {
		req.SetBasicAuth(username, password)
	}

	// https://www.cnblogs.com/gaorong/p/11336834.html
	// https://blog.cloudflare.com/the-complete-guide-to-golang-net-http-timeouts/
	cli := &http.Client{
		Timeout: time.Duration(timeoutSeconds) * time.Second,
	}
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
