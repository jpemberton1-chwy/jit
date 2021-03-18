package app

import (
	"fmt"
	"strings"

	"github.com/jpemberton1-chwy/jit/git"
)

func RunLog(args []string) (success bool, err error) {
	history, err := git.ReadCommitHistory()

	if err != nil {
		return false, err
	}

	for _, line := range strings.Split(history, "\n") {
		if len(line) < 82 {
			break
		}
		commit := git.CreateCommitFromLog(line)
		fmt.Println(commit.Message)
	}

	return true, nil
}
