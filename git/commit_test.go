package git

import (
	"log"
	"os"
	"testing"

	. "github.com/franela/goblin"
)

func TestCommit(t *testing.T) {
	g := Goblin(t)
	g.Describe("commit", func() {
		g.AfterEach(func() {
			os.RemoveAll("/tmp/.git")
		})

		g.Describe("with git repository and commits", func() {
			g.BeforeEach(func() {
				os.Chdir("/tmp")
				os.Mkdir(".git", 0700)
				os.Mkdir(".git/logs", 0700)
				f, err := os.OpenFile(".git/logs/HEAD", os.O_RDWR|os.O_CREATE, 0700)
				if err != nil {
					log.Fatal("failed to open file")
				}
				defer f.Close()

				commits := "0000000000000000000000000000000000000000 9e0554231d89892865420a4058f1529165888374 Guy Pierce <gp@example.com> 1615582519 -0500  commit (initial): Initial commit"
				f.Write([]byte(commits))
			})

			g.It("should parse git history", func() {
				res, err := ReadCommitHistory()
				g.Assert(err).IsNil()
				g.Assert(res).IsNotNil()
			})
		})

		g.Describe("with no git repository", func() {
			g.BeforeEach(func() {
				os.Chdir("/tmp")
			})

			g.It("should fail to parse commit log if not in repository", func() {
				_, err := ReadCommitHistory()
				g.Assert(err).IsNotNil()
			})
		})

		g.Describe("when read error occurs", func() {
			g.BeforeEach(func() {
				os.Chdir("/tmp")
				os.Mkdir(".git", 0700)
				os.Mkdir(".git/logs", 0700)
				os.Create(".git/logs/HEAD")
				os.Chmod(".git/logs/HEAD", 0000)
			})

			g.It("should fail to parse commit log if error", func() {
				_, err := ReadCommitHistory()
				g.Assert(err).IsNotNil()
			})
		})
	})
}
