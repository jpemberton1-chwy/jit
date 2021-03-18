package args

import (
	"testing"

	"github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func TestParser(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("ParseArgs", func() {
		g.It("should parse a command with no args", func() {
			args := ParseArgs("command", []string{})
			Expect(args.Command).To(Equal("command"))
			Expect(args.GetArgs()).To(BeEmpty())
		})

		g.It("should parse a command with args", func() {
			args := ParseArgs("command", []string{"arg1"})
			Expect(args.GetArgs()).To(HaveLen(1))
			Expect(args.GetArgs()).To(ContainElement("arg1"))
		})

		g.It("should parse a command with options", func() {
			args := ParseArgs("command", []string{"--option", "value"})
			Expect(args.GetOption("option")).To(Equal("value"))
		})

		g.It("should parse a command with options and arguments", func() {
			args := ParseArgs("command", []string{"arg1", "--option", "value", "arg2"})
			Expect(args.HasArg("arg1")).To(BeTrue())
			Expect(args.HasArg("arg2")).To(BeTrue())
			Expect(args.GetOption("option")).To(Equal("value"))
		})
	})
}
