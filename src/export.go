package main

func init() {
	export := CommandOptions{
		name:    "export",
		summary: "export all variables in an environment",
		usage:   "[<options>] export [<args>]",
	}
	commands = append(commands, export)
}
