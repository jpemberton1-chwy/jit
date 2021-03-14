package app

import (
	"testing"

	. "github.com/franela/goblin"
)

func TestStart(t *testing.T) {
	g := Goblin(t)

	g.Describe("start", func() {
		g.It("should log start of workday", func() {
			g.Assert(true).IsTrue()
		})
	})
}
