package app

import (
	"testing"

	. "github.com/franela/goblin"
)

func Test(t *testing.T) {
	g := Goblin(t)
	g.Describe("ParseArgs", func() {
		g.It("should parse log subcommand", func() {
			args, _ := ParseArgs([]string{"log"})
			g.Assert(args.Subcommand).Equal("log")
		})
		g.It("should parse log subcommand args", func() {
			args, _ := ParseArgs([]string{"log", "CSRB-50"})
			g.Assert(args.SubcommandArgs).Equal([]string{"CSRB-50"})
		})
	})
}
