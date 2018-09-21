package main

func init() {
	add := CommandOptions{
		name:    "add",
		summary: "add a new variable to an environment",
		usage:   "[<options>] add [<args>]",
	}
	commands = append(commands, add)
}
