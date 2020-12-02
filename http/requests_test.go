package http

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"
)

func Test_Request(t *testing.T) {
	req := NewRequest("https://www.baidu.com")
	//req.SetUri("")
	fmt.Println(req.GetURL())
	reqErr := req.Request(nil)
	if reqErr != nil {
		fmt.Println(reqErr)
	}

	resErr := req.Do()
	if resErr != nil {
		fmt.Println(resErr)
	}

	defer req.Response.Body.Close()
	body, err := ioutil.ReadAll(req.Response.Body)
	if err != nil {
		fmt.Println("http 读取响应失败")
	}

	fmt.Println(string(body))

	fmt.Println(req.Response)
	fmt.Println(req.Response.Status)
	fmt.Println(req.Response.Proto)

	fmt.Println(req.Response.ContentLength)
	fmt.Println(req.Response.Request)

}

func Test_Post(t *testing.T) {
	req := NewRequest("http://data-metadata.c.goops.top/user/getUserInfo")
	//req.SetUri("")
	fmt.Println(req.GetURL())
	// name/departmentName/email/mobile 是一个动态参数
	respErr := req.Post(bytes.NewReader([]byte(`{"username":"dev_ops","password":"nqJTgY4Z","mobile":"1xxxxxxx"}`)))
	if respErr != nil {
		fmt.Println(respErr)
	}

	defer req.Response.Body.Close()
	body, err := ioutil.ReadAll(req.Response.Body)
	if err != nil {
		fmt.Println("http 读取响应失败")
	}

	fmt.Println(string(body))

}

func Test_Head(t *testing.T) {
	req := NewRequest("http://data-metadata.c.goops.top/user/getUserInfo")
	headErr := req.Head()

	if headErr != nil {
		fmt.Println(headErr)
	}

	fmt.Println(req.Response.Header)

}
