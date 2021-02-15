package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	for {
		resp, err := http.Get("http://hits.dwyl.com/complexorganizations/git-views-generator.svg")
		if err != nil {
			log.Println(err)
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(string(body))
	}
}
