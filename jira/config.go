package jira

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type JiraConfig struct {
	Auth string
}

func Load() (res JiraConfig, err error) {
	dir, err := os.UserHomeDir()
	if err != nil {
		return JiraConfig{}, err
	}

	jitrc, err := os.Open(fmt.Sprintf("%s/.jitrc", dir))
	if err != nil {
		return JiraConfig{}, err
	}
	defer jitrc.Close()

	bytes, err := ioutil.ReadAll(jitrc)
	if err != nil {
		return JiraConfig{}, err
	}

	var jc JiraConfig
	json.Unmarshal(bytes, &jc)
	return jc, nil
}

func Create(username string, apiKey string) JiraConfig {
	auth := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", username, apiKey)))
	return JiraConfig{
		Auth: auth,
	}
}

func Save(config JiraConfig) (success bool, err error) {
	dir, err := os.UserHomeDir()
	if err != nil {
		return false, err
	}

	file, err := os.OpenFile(fmt.Sprintf("%s/.jitrc", dir), os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		return false, err
	}
	defer file.Close()

	json, err := json.Marshal(config)
	if err != nil {
		return false, err
	}

	file.Write(json)

	return true, nil
}
