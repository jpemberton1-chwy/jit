package app

import "errors"

type JitArgs struct {
	Subcommand     string
	SubcommandArgs []string
}

func ParseArgs(programArgs []string) (result JitArgs, err error) {
	var args JitArgs = JitArgs{}
	var subcommand string = "help"
	if len(programArgs) > 0 {
		subcommand = programArgs[0]
	}

	if isValidSubcommand(subcommand) {
		args.Subcommand = subcommand
		args.SubcommandArgs = append(args.SubcommandArgs, programArgs[1:]...)
	} else {
		return args, errors.New("failed to parse arguments")
	}

	return args, nil
}

var subcommands []string = []string{
	"log",
	"start",
	"stop",
	"help",
}

func isValidSubcommand(subcommand string) bool {
	for i := 0; i < len(subcommands); i += 1 {
		if subcommands[i] == subcommand {
			return true
		}
	}
	return false
}
