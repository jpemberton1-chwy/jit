package app

import (
	"fmt"
	"os"
	"time"
)

func RunStart(args []string) (success bool, err error) {
	dir, err := os.UserHomeDir()
	if err != nil {
		return false, err
	}

	file, err := os.OpenFile(fmt.Sprintf("%s/.jit/starttime", dir), os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0600)
	if err != nil {
		return false, err
	}
	defer file.Close()

	file.WriteString(fmt.Sprintf("%d", time.Now().Unix()))

	return true, nil
}
