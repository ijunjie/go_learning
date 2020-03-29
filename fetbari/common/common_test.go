package common

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOnlineOrOffline(t *testing.T) {
	online := OnlineOrOffline("online")
	t.Logf("env: %d", online)
}

func TestBase64Encode(t *testing.T) {
	encode := Base64Encode("admin", "admin")
	assert.Equal(t, "YWRtaW46YWRtaW4=", encode)
}


func TestHttpGet(t *testing.T) {
	body, _ := HttpGet("http://httpbin.org/json", "", "")
	var res Resp
	_ = json.Unmarshal(body, &res)
	fmt.Printf("%+v", res.SlideShow)
}

type Resp struct {
	SlideShow SlideShow
}

type SlideShow struct {
	Author string `json:"author"`
	Title string `json:"title"`

}