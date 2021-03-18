package app

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

const author = "Jonathon Pemberton"
const email string = "jpemberton1@chewy.com"
const program string = "jit"
const version string = "v1.0.0"
const year = "2021"
const terminalChars int = 60
const commandPadding int = 4

func RunHelp() {
	version := fmt.Sprintf("%s %s - %s <%s>", program, version, author, email)
	color.Cyan(version)
	color.Cyan(fmt.Sprintf("           ©️ %s %s", year, author))
	color.Cyan("")
	color.Cyan("Time tracking through Git 🚀")
	color.Cyan("")
	color.Cyan("Available commands")
	PrintCmd("log", "Rollup the history for each commit")
	PrintCmd("help", "Rollup the history for each commit")
}

func PrintCmd(command string, help string) {
	padding := []string{}
	for i := 0; i < commandPadding; i += 1 {
		padding = append(padding, " ")
	}
	spaces := []string{}
	for i := 0; i < terminalChars-len(command)-len(help)-commandPadding; i += 1 {
		spaces = append(spaces, ".")
	}
	color.Cyan(fmt.Sprintf("%s%s%s%s", strings.Join(padding, ""), command, strings.Join(spaces, ""), help))
}
