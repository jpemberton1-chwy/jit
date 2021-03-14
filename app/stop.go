package app

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

func RunStop(args []string) (success bool, err error) {
	dir, err := os.UserHomeDir()
	if err != nil {
		return false, err
	}

	file, err := os.OpenFile(fmt.Sprintf("%s/.jit/starttime", dir), os.O_RDONLY, 0600)
	if err != nil {
		return false, err
	}
	defer file.Close()

	startTimeBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return false, err
	}

	startTimeDigits, err := strconv.ParseInt(string(startTimeBytes[:]), 10, 64)
	if err != nil {
		return false, err
	}

	startTime := time.Unix(startTimeDigits, 0)
	fmt.Println(startTime)
	fmt.Printf("%s\n", time.Since(startTime))
	return true, nil
}
