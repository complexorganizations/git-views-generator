package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	for {
		resp, err := http.Get("https://hits.seeyoufarm.com/api/count/incr/badge.svg?url=https%3A%2F%2Fgithub.com%2Fcomplexorganizations%2Fgit-views-generator&count_bg=%2379C83D&title_bg=%23555555&icon=&icon_color=%23E7E7E7&title=hits&edge_flat=false")
		if err != nil {
			log.Println(err)
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(string(body))
	}
}
