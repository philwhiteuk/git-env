package commands

import (
	"bytes"
	"flag"
	"fmt"
	"strings"
)

// Command represents the public interface for a command in git-env
type Command interface {
	Name() string
	Summary() string
	Usage() bytes.Buffer
	Run() error
}

func commands() []Command {
	var c []Command
	return append(c,
		Add,
		Create,
		Export,
		Help,
		List,
		Rotate,
		Subst,
	)
}

// Get returns a single instance of a Command in git-env
// returns an error if no such command exists
func Get(r string) (Command, error) {
	for _, co := range commands() {
		if co.Name() == r {
			return co, nil
		}
	}

	return nil, fmt.Errorf("'%s' is not a git-env command", r)
}

type commandOptions struct {
	flags   *flag.FlagSet
	name    string
	summary string
	usage   string
}

// Name returns the name of the command
func (c commandOptions) Name() string {
	return c.name
}

// Summary returns a summary of what the command does
func (c commandOptions) Summary() string {
	return c.summary
}

// Usage returns usage information about the command
func (c commandOptions) Usage() bytes.Buffer {
	var b bytes.Buffer

	b.WriteString(fmt.Sprintf("usage: git env %s\n", c.usage))
	b.WriteString(c.flagUsage())

	return b
}

func (c commandOptions) flagUsage() string {
	var b bytes.Buffer

	if c.flags != nil {
		b.WriteString(fmt.Sprintf("\n"))

		defaults := func(f *flag.Flag) {
			if f.Usage != "" {
				prefix := string([]rune(f.Name)[1:2])

				var flagName string
				shortOption := c.flags.Lookup(prefix)
				if shortOption != nil && shortOption.Usage == "" {
					var n []string
					n = append(n, prefix, f.Name)
					flagName = strings.Join(n, " -")
				} else {
					flagName = f.Name
				}
				b.WriteString(fmt.Sprintf("    -%-20s %-14s\n", flagName, f.Usage))
			}
		}

		c.flags.VisitAll(defaults)
	}

	return b.String()
}
