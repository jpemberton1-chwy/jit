package args

import "github.com/jpemberton1-chwy/jit/util"

type JitArgs struct {
	Command string
	Args    *util.Set
	Options map[string]interface{}
}

func NewArgs(command string) *JitArgs {
	return &JitArgs{
		Command: command,
		Args:    util.NewSet(),
		Options: map[string]interface{}{},
	}
}

func (a *JitArgs) AddArg(arg string) *JitArgs {
	a.Args.Add(arg)
	return a
}

func (a *JitArgs) AddOption(option string, value interface{}) *JitArgs {
	a.Options[option] = value
	return a
}

func (a *JitArgs) HasArg(arg string) bool {
	return a.Args.Contains(arg)
}

func (a *JitArgs) HasOption(option string) bool {
	_, c := a.Options[option]
	return c
}

func (a *JitArgs) GetOption(option string) interface{} {
	element := a.Options[option]
	return element
}

func (a *JitArgs) GetArgs() []interface{} {
	return a.Args.AsSlice()
}
