package main

import (
	"io/ioutil"
	"net/http"
	"time"
)

func main() {

	c := http.Client{Timeout: time.Nanosecond}
	resp, err := c.Get("http://www.google.com")
	if err != err {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	println(string(body))
}
