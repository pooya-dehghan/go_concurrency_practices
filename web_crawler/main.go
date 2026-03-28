package main

import (
	"fmt"
	"net/http"
)

//Objective : we want to crawl all the links in a webpage and have a depth parameter to determine how deep we want to crawl

func main() {
	link := "https://isna.ir"

	var links []string

	res, err := http.Get(link)
	defer res.Body.Close()
	if err != nil {
		fmt.Printf("there was some error in getting link specified %s and error is %v ", link, err)
	}

	http.parse

}
