package client

import (
	"beego/config"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func Get(c chan int, i int)  {
	host := config.HOST
	url := host + config.API + "?user=go"

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

	c <- i

	fmt.Println(string(body))
}

func Post(c chan int, i int)  {
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

	c <- i

	fmt.Println(string(body))
}


