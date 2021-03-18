package main

import "fmt"

func main() {
	fmt.Println("hello world")
	// jitArgs := &JitArgs{
	// 	Command: os.Args[1],
	// 	Args:    NewSet(),
	// 	Options: NewSet(),
	// }

	// for _, arg := range os.Args[1:] {
	// 	if strings.HasPrefix(arg, "--") {
	// 		jitArgs.Options.Add(arg[2:])
	// 	} else {
	// 		jitArgs.Args.Add(arg)
	// 	}
	// }
}
