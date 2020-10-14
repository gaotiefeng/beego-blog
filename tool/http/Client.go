package http

import (
	"beego/app/constants"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func Get()  {
	host := constants.HOST
	url := host + constants.API + "?user=go"

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
	host := constants.HOST
	url := host + constants.API

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


