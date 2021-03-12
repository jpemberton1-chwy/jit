package jira

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func FetchData() {
	resp, err := http.Get("https://google.com")
	if err != nil {
		log.Fatal("fatal")
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("fatal")
	}
	fmt.Println(string(content[:]))
}
