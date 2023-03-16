package client

import (
	"beego/config"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func Get() {
	host := config.HOST
	url := host + config.API + "?user=go"

	resp, err := http.Get(url)
	if err != nil {
		// handle error
		fmt.Println("client error", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println("client get out", string(body))
}

func Post() {
	host := config.HOST
	url := host + config.API

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

	fmt.Println("client post out", string(body))
}
