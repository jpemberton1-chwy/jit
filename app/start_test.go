package app

import (
	"fmt"
	"os"
	"testing"

	. "github.com/franela/goblin"
)

func TestStart(t *testing.T) {
	g := Goblin(t)

	g.Describe("start", func() {
		g.BeforeEach(func() {
			dir, err := os.UserHomeDir()
			if err != nil {
				panic("can't get user home")
			}
			os.Mkdir(fmt.Sprintf("%s/.jit", dir), 0600)
		})

		g.AfterEach(func() {
			dir, err := os.UserHomeDir()
			if err != nil {
				panic("can't get user home")
			}
			os.RemoveAll(fmt.Sprintf("%s/.jit", dir))
		})

		g.It("should log start of workday", func() {
			res, err := RunStart([]string{})
			g.Assert(err).IsNil()
			g.Assert(res).Equal(true)
		})
	})
}
