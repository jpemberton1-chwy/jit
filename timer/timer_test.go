package timer

import (
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/franela/goblin"
	"github.com/jpemberton1-chwy/jit/util"
	. "github.com/onsi/gomega"
)

func TestTimer(t *testing.T) {
	g := goblin.Goblin(t)

	RegisterFailHandler(func(m string, _ ...int) {
		g.Fail(m)
	})

	g.Describe("timer", func() {
		g.Describe("Start", func() {
			g.BeforeEach(func() {
				os.Mkdir(util.GetUserFilePath(".jit"), 0755)
				util.DeleteTimerFile()
				// if err != nil {
				// 	panic(fmt.Sprintf("can't remove file %s", util.GetUserFilePath(".jit/timer")))
				// }
			})

			g.AfterEach(func() {
				os.Remove(util.GetUserFilePath(".jit"))
			})

			g.It("should create a timer file", func() {
				res, err := Start(time.Now())
				Expect(err).ShouldNot(HaveOccurred())
				Expect(res).To(Equal(true))
			})

			g.It("should log specified start time", func() {
				_, err := Start(time.Date(2020, time.December, 25, 6, 25, 20, 0, time.UTC))
				Expect(err).ShouldNot(HaveOccurred())
				file, err := util.OpenUserFile(".jit/timer")
				Expect(err).ShouldNot(HaveOccurred())
				content, err := ioutil.ReadAll(file)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(string(content[:])).To(Equal("2020-12-25 06:25:20"))
			})

			g.Describe("if timer has already been started", func() {
				g.It("should return error", func() {
					_, err := Start(time.Now())
					Expect(err).ShouldNot(HaveOccurred())
					_, err = Start(time.Now())
					Expect(err).Should(HaveOccurred())
				})
			})
		})
	})
}
