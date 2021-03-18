package util

import (
	"fmt"
	"os"
)

// Holds various cacheable values
var cache map[string]string = map[string]string{
	"home": "",
}

func GetUserHomeDir() string {
	if cache["home"] == "" {
		res, err := os.UserHomeDir()
		if err != nil {
			panic("failed to get user home")
		}
		cache["home"] = res
	}

	return cache["home"]
}

func GetUserFilePath(filename string) string {
	return fmt.Sprintf("%s/%s", GetUserHomeDir(), filename)
}

func OpenUserFile(filename string) (result *os.File, err error) {
	home := GetUserHomeDir()
	path := fmt.Sprintf("%s/%s", home, filename)

	result, err = os.OpenFile(path, os.O_RDONLY, 0600)
	if err != nil {
		return &os.File{}, err
	}

	return result, nil
}

func FileExists(filename string) bool {
	file, err := os.Open(filename)
	if err != nil {
		return false
	}
	defer file.Close()
	return true
}

func DeleteTimerFile() error {
	return os.Remove(GetUserFilePath(".jit/timer"))
}
