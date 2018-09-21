package main

var help CommandOptions

func init() {
	help = CommandOptions{
		name:    "help",
		summary: "display information about a command",
		usage:   "help [<command>]",
	}
	commands = append(commands, help)
}
