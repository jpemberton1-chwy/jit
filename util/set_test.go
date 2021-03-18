package util

import (
	"reflect"
	"testing"

	"github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func TestSet(t *testing.T) {
	g := goblin.Goblin(t)

	RegisterFailHandler(func(m string, _ ...int) {
		g.Fail(m)
	})

	g.Describe("set", func() {
		g.It("should create a new set", func() {
			s := NewSet()
			Expect(reflect.TypeOf(s).String()).To(Equal("*util.set"))
		})

		g.Describe("Add", func() {
			g.It("should add item to set", func() {
				s := NewSet().Add("thing")
				Expect(s.m["thing"]).To(Equal(exists))
			})
		})

		g.Describe("Contains", func() {
			g.It("should report if item is in set", func() {
				s := NewSet().Add("thing")
				Expect(s.Contains("thing")).To(BeTrue())
			})

			g.It("should report if item is not in set", func() {
				s := NewSet()
				Expect(s.Contains("thing")).To(BeFalse())
			})
		})

		g.Describe("Remove", func() {
			g.It("should remove item from set", func() {
				s := NewSet().Add("thing").Remove("thing")
				Expect(s.Contains("thing")).To(BeFalse())
			})
		})
	})
}
