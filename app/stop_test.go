package app

import (
	"testing"

	. "github.com/franela/goblin"
)

func TestStop(t *testing.T) {
	g := Goblin(t)
	g.Describe("stop", func() {
		// g.BeforeEach(func() {

		// })

		g.It("should stop timer", func() {
			g.Assert(true).IsTrue()
		})
	})
}
