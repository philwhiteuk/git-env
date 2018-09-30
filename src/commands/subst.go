package commands

type substCommand struct {
	commandOptions
}

// Subst subst command
var Subst Command

func init() {
	Subst = substCommand{
		commandOptions{
			name:    "subst",
			summary: "substitute file contents with environment variables",
			usage:   "[<options>] subst [<args>]",
		},
	}
}

func (h substCommand) Run() error {
	return nil
}
