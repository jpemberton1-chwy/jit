package test

import (
	"io/ioutil"
	"os"
)

func Cap(f func()) string {
	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = stdOut

	return string(out[:])
}
