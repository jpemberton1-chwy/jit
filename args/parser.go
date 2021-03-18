package args

import "strings"

func ParseArgs(command string, arguments []string) *JitArgs {
	args := NewArgs(command)
	wasOption := false
	var key string
	for _, arg := range arguments {
		if strings.HasPrefix(arg, "--") {
			key = arg[2:]
			wasOption = true
		} else if wasOption {
			args.AddOption(key, arg)
			wasOption = false
		} else {
			args.AddArg(arg)
		}
	}
	return args
}
