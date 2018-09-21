package main

func init() {
	subst := CommandOptions{
		name:    "subst",
		summary: "substitute file contents with environment variables",
		usage:   "[<options>] subst [<args>]",
	}
	commands = append(commands, subst)
}
