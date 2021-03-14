package jira

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type JiraConfig struct {
	Auth string
}

func Load() JiraConfig {
	file, err := os.Open("/Users/jpemberton1/.jitrc")
	if err != nil {
		log.Fatal(err)
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	var jc JiraConfig
	json.Unmarshal(bytes, &jc)
	return jc
}

func Create(username string, apiKey string) JiraConfig {
	auth := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", username, apiKey)))
	return JiraConfig{
		Auth: auth,
	}
}

func Save(config JiraConfig) {
	file, err := os.OpenFile("/Users/jpemberton1/.jitrc", os.O_CREATE|os.O_RDWR, 0700)
	if err != nil {
		file.Close()
		log.Fatal(err)
	}
	json, err := json.Marshal(config)
	if err != nil {
		file.Close()
		log.Fatal(err)
	}
	file.Write(json)
}
