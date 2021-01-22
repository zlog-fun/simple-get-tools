package myhttp

import (
	"fmt"
	"net/http"
)

// URLStructInfo url请求信息
type URLStructInfo struct {
	Url  string
	Host string
	Port string
	Path string
}

// Get 发起get请求
func (r *URLStructInfo) Get() bool {
	//fmt.Println(r.host)
	pre := "http://"
	if r.Port == "443" {
		pre = "https://"
	}
	url := fmt.Sprintf("%s%s/%s", pre, r.Host, r.Path)
	if r.Url != "" {
		url = fmt.Sprintf("%s%s/%s", pre, r.Url, r.Path)
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	if r.Url != "" {
		req.Host = r.Host
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	// fmt.Println(r.Url, resp.StatusCode, resp.ContentLength)
	if resp.StatusCode != 200 {
		return false
	}
	return true
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	panic(err)
	// }
}
