package jira

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func FetchData(config JiraConfig) {
	client := http.Client{}
	req, err := http.NewRequest("GET", "https://chewyinc.atlassian.net/rest/api/latest/issue/CSRB-50", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", config.Auth)
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content[:]))
}
