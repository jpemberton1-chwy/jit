package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

type JitArgs struct {
	Subcommand     string
	SubcommandArgs []string
}

const author = "Jonathon Pemberton"
const email string = "jpemberton1@chewy.com"
const program string = "jit"
const version string = "v1.0.0"
const year = "2021"
const terminalChars int = 60
const commandPadding int = 4

var subcommands []string = []string{
	"log",
}

func main() {
	if len(os.Args) < 2 {
		PrintUsage()
		os.Exit(0)
	}

	subcommand := os.Args[1]

	if IsValidSubcommand(subcommand) {
		jitArgs := JitArgs{Subcommand: subcommand}
		jitArgs.SubcommandArgs = append(jitArgs.SubcommandArgs, os.Args[2:]...)

		HandleArgs(jitArgs)
	} else {
		color.Red(fmt.Sprintf("jit: Invalid argument %s", subcommand))
		os.Exit(2)
	}
}

func IsValidSubcommand(subcommand string) bool {
	for i := 0; i < len(subcommands); i += 1 {
		if subcommands[i] == subcommand {
			return true
		}
	}
	return false
}

func PrintUsage() {
	version := fmt.Sprintf("%s %s - %s <%s>", program, version, author, email)
	color.Cyan(version)
	color.Cyan(fmt.Sprintf("           ©️ %s %s", year, author))
	color.Cyan("")
	color.Cyan("Time tracking through Git 🚀")
	color.Cyan("")
	color.Cyan("Available commands")
	PrintCmd("log", "Rollup the history for each commit")
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

func HandleArgs(jitArgs JitArgs) {
	var f func(args []string)
	if jitArgs.Subcommand == "log" {
		f = HandleLog
	}
	f(jitArgs.SubcommandArgs)
}

func HandleLog(args []string) {

}
