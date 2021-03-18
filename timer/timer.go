package timer

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/jpemberton1-chwy/jit/util"
)

func Start(startTime time.Time) (success bool, err error) {
	file, err := os.OpenFile(util.GetUserFilePath(".jit/timer"), os.O_CREATE|os.O_EXCL, 0600)
	if err != nil {
		return false, err
	}
	defer file.Close()

	_, err = file.WriteString(startTime.Format("2006-01-02 15:04:05"))
	if err != nil {
		return false, err
	}

	return true, nil
}

func Stop() (success bool, err error) {
	file, err := os.Open(util.GetUserFilePath(".jit/timer"))
	if err != nil {
		return false, err
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		return false, err
	}

	startTime, err := time.Parse("2006-01-02 03:04:05", string(content[:]))
	if err != nil {
		return false, err
	}

	value := time.Since(startTime)
	fmt.Println(value)
	return true, nil
}
