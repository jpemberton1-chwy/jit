package util

import (
	"os"
	"testing"

	"github.com/franela/goblin"
)

func TestUtil(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("util", func() {
		g.Describe("GetUserHomeDir", func() {
			var expected string = ""
			g.BeforeEach(func() {
				res, err := os.UserHomeDir()
				if err != nil {
					panic("failed to get user home")
				}
				expected = res
			})

			g.It("should get user home directory", func() {
				home := GetUserHomeDir()
				g.Assert(home).Equal(expected)
			})

			g.It("should cache user home directory for subsequent calls", func() {
				home := GetUserHomeDir()
				g.Assert(GetUserHomeDir()).Equal(home)
			})
		})
	})
}
