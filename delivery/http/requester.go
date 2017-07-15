package main

import "net/http"
import "io/ioutil"
import "fmt"

type Http interface {
}

func RequestGoogle() []byte {
	resp, err := http.Get("http://www.google.com")
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("SHITS BROKE")
		panic(err.Error())
	}

	return body
}
