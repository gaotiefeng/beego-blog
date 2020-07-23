package client

import (
	"beego/conf"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func Get()  {
	host := conf.HOST
	url := host + conf.API + "?user=go"

	resp, err := http.Get(url)
	if err != nil {
		// handle error
		fmt.Println("client error",err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println("client get out",string(body))
}

func Post()  {
	host := conf.HOST
	url := host + conf.API

	resp, err := http.Post(url,
		"application/x-www-form-urlencoded",
		strings.NewReader("user=go"))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println("client post out",string(body))
}


