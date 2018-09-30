package commands

type addCommand struct {
	commandOptions
}

// Add add command
var Add Command

func init() {
	Add = addCommand{
		commandOptions{
			name:    "add",
			summary: "add a new variable to an environment",
			usage:   "[<options>] add [<args>]",
		},
	}
}

func (h addCommand) Run() error {
	return nil
}
