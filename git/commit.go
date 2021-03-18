package git

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"
)

type Commit struct {
	ParentCommitHash   string
	CommitHash         string
	Author             string
	Email              string
	TimestampUnixEpoch string
	TimestampLocale    string
	Message            string
}

func CreateCommitFromLog(log string) Commit {
	if len(log) < 82 {
		panic("failed to parse commit")
	}

	commit := Commit{
		ParentCommitHash: log[0:40],
		CommitHash:       log[41:81],
	}

	var authorEnd int = 0
	for j, char := range []rune(log[82:]) {
		if char == 62 {
			authorEnd = j + 83
			break
		}
	}
	commit.Author = strings.Trim(log[82:authorEnd], " \t")

	var timestampEpochEnd int = 0
	for j, char := range []rune(log[authorEnd+1:]) {
		if char == 32 {
			timestampEpochEnd = j + authorEnd
			break
		}
	}
	commit.TimestampUnixEpoch = strings.Trim(log[authorEnd:timestampEpochEnd], " \t")

	var localeEnd int = 0
	for j, char := range []rune(log[timestampEpochEnd+2:]) {
		if char == 9 {
			localeEnd = j + timestampEpochEnd + 2
			break
		}
	}
	commit.TimestampLocale = strings.Trim(log[timestampEpochEnd+2:localeEnd], " \t")
	commit.Message = strings.Trim(log[localeEnd:], " \t")

	return commit
}

func ReadCommitHistory() (result string, err error) {
	file, err := os.Open(".git/logs/HEAD")
	if err != nil {
		return "", errors.New("failed to read commit history, are you in a git repository?")
	}

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		file.Close()
		return "", errors.New("an unexpected read error has occurred")
	}

	return string(bytes[:]), nil
}
