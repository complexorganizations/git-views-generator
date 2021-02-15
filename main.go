package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	for {
		resp, err := http.Get("https://visitor-badge.glitch.me/badge?page_id=complexorganizations.git-views-generator")
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
