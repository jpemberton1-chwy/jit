package args

import (
	"testing"

	"github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func TestArgs(t *testing.T) {
	g := goblin.Goblin(t)

	RegisterFailHandler(func(m string, _ ...int) {
		g.Fail(m)
	})

	g.Describe("JitArgs", func() {
		g.It("should create new args struct", func() {
			args := NewArgs("thing")
			Expect(args.Command).To(Equal("thing"))
		})

		g.It("should add args", func() {
			args := NewArgs("thing")
			args.AddArg("argument")
			Expect(args.HasArg("argument")).To(BeTrue())
		})

		g.It("should retrieve args", func() {
			args := NewArgs("thing").AddArg("one").AddArg("two").AddArg("three")
			Expect(args.GetArgs()).To(ContainElements("one", "two", "three"))
		})

		g.It("should add options", func() {
			args := NewArgs("thing")
			args.AddOption("option", 123)
			Expect(args.HasOption("option")).To(BeTrue())
		})

		g.It("should retrieve options", func() {
			args := NewArgs("thing")
			args.AddOption("option", 123)
			Expect(args.GetOption("option")).To(Equal(123))
		})
	})
}
