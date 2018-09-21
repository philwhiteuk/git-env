package main

func init() {
	rotate := CommandOptions{
		name:    "rotate",
		summary: "rotate an existing environment variable",
		usage:   "[<options>] rotate [<args>]",
	}
	commands = append(commands, rotate)
}
