package main

import (
	"commands"
	"flag"
)

func main() {
	flag.Parse()

	if flag.NArg() > 0 {
		command, err := commands.Get(flag.Arg(0))
		if err == nil {
			command.Run()
			return
		}
	}

	commands.Help.Run()
}
