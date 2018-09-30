package commands

type rotateCommand struct {
	commandOptions
}

// Rotate rotate command
var Rotate Command

func init() {
	Rotate = rotateCommand{
		commandOptions{
			name:    "rotate",
			summary: "rotate an existing environment variable",
			usage:   "[<options>] rotate [<args>]",
		},
	}
}

func (h rotateCommand) Run() error {
	return nil
}
