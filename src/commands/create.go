package commands

type createCommand struct {
	commandOptions
}

// Create create command
var Create Command

func init() {
	Create = createCommand{
		commandOptions{
			name:    "create",
			summary: "create a new environment",
			usage:   "[<options>] create [<args>]",
		},
	}
}

func (h createCommand) Run() error {
	return nil
}
