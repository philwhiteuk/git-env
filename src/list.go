package main

import (
	"flag"
)

func init() {
	f := flag.NewFlagSet("list", flag.ContinueOnError)
	f.Bool("all", false, "shows all environments")

	list := CommandOptions{
		flags:   f,
		name:    "list",
		summary: "list all variables stored in an environment",
		usage:   "[<options>] list [<args>]",
	}
	commands = append(commands, list)
}
