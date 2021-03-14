package jira_test

import (
	"encoding/base64"
	"os"
	"testing"

	"github.com/jpemberton1-chwy/jit/jira"
)

func TestCreateShouldCreateJiraConfig(t *testing.T) {
	expected := base64.StdEncoding.EncodeToString([]byte("jpemberton1@chewy.com:1234"))
	config := jira.Create("jpemberton1@chewy.com", "1234")
	if config.Auth != expected {
		t.Error("failed to create JiraConfig with expected auth string")
	}
}

func TestSaveShouldSaveJitrcFile(t *testing.T) {
	jira.Save(jira.Create("jpemberton1@chewy.com", "1234"))
	_, err := os.OpenFile("/Users/jpemberton1/.jitrc", os.O_RDONLY, 0000)
	if err != nil {
		t.Error("did not create .jitrc")
	}
}

func TestLoadShouldLoadJitrcFile(t *testing.T) {
	config := jira.Load()
	if config.Auth == "" {
		t.Error("failed to load .jitrc")
	}
}
