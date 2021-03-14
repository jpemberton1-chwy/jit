package jira_test

import (
	"testing"

	. "github.com/franela/goblin"
	"github.com/jpemberton1-chwy/jit/jira"
)

func TestConfig(t *testing.T) {
	g := Goblin(t)
	g.Describe("config", func() {
		g.Describe("load", func() {
			g.Describe("when configuration exists", func() {
				g.BeforeEach(func() {
					jira.Save(jira.Create("jpemberton1@example.com", "1234"))
				})

				g.It("should load .jitrc", func() {
					config, err := jira.Load()
					g.Assert(err).IsNil()
					g.Assert(config.Auth).Equal("anBlbWJlcnRvbjFAZXhhbXBsZS5jb206MTIzNA==")
				})
			})
		})
	})
}
