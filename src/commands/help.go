package commands

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
)

type helpCommand struct {
	commandOptions
}

// Help help command
var Help Command

func init() {
	flag.Usage = func() {
		Help.Run()
	}
	flag.CommandLine.SetOutput(os.Stdout)
	flag.CommandLine.Init("defaults", flag.ContinueOnError)

	Help = helpCommand{
		commandOptions{
			name:    "help",
			summary: "display information about a command",
			usage:   "help [<command>]",
		},
	}
}

func (h helpCommand) Run() error {
	var command Command
	var err error
	if flag.NArg() == 1 {
		command, err = Get(flag.Arg(0))
	} else if flag.NArg() > 1 {
		command, err = Get(flag.Arg(1))
	}

	var usage bytes.Buffer
	if command != nil {
		usage = command.Usage()
	} else {
		usage = defaultUsage()
		if err != nil {
			l := log.New(os.Stderr, "", 0)
			l.Println(fmt.Sprintf("git-env: %s. See 'git env --help'.", err))

			os.Exit(2)
		}
	}

	l := log.New(os.Stdout, "", 0)
	l.Println(usage.String())

	return nil
}

func defaultUsage() bytes.Buffer {
	defaults := commandOptions{
		usage: "[-h | --help] [<command>] [<args>]",
	}
	b := defaults.Usage()

	b.WriteString(fmt.Sprintf("store and encrypt environment variables natively in git\n\n"))
	for _, c := range commands() {
		if c != nil {
			b.WriteString(fmt.Sprintf("   %-11s %-14s\n", c.Name(), c.Summary()))
		}
	}

	return b
}
