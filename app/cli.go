package app

import (
	"os"

	"github.com/fatih/color"
)

func Run() {
	args, err := ParseArgs(os.Args[1:])

	if err == nil {
		RunCmd(args)
	} else {
		color.Red("jit: %s", err.Error())
	}
}

func RunCmd(args JitArgs) {
	var f func(args []string) (success bool, err error)

	if args.Subcommand == "log" {
		f = RunLog
	} else if args.Subcommand == "start" {
		f = RunStart
	} else if args.Subcommand == "stop" {
		f = RunStop
	}

	_, err := f(args.SubcommandArgs)
	if err != nil {
		color.Red("jit: %s", err.Error())
	}
}
