package commands

type exportCommand struct {
	commandOptions
}

// Export export command
var Export Command

func init() {
	Export = exportCommand{
		commandOptions{
			name:    "export",
			summary: "export all variables in an environment",
			usage:   "[<options>] export [<args>]",
		},
	}
}

func (h exportCommand) Run() error {
	return nil
}
