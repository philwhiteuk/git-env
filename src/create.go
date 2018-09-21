package main

func init() {
	create := CommandOptions{
		name:    "create",
		summary: "create a new environment",
		usage:   "[<options>] create [<args>]",
	}
	commands = append(commands, create)
}
