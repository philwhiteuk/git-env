package commands

import (
	"flag"
)

type listCommand struct {
	commandOptions
}

// List list command
var List Command

func init() {
	f := flag.NewFlagSet("list", flag.ContinueOnError)
	f.Bool("a", false, "")
	f.Bool("-all", false, "shows all environments")

	List = listCommand{
		commandOptions{
			flags:   f,
			name:    "list",
			summary: "list all variables stored in an environment",
			usage:   "[<options>] list [<args>]",
		},
	}
}

func (h listCommand) Run() error {
	return nil
}
