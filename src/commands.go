package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type Command interface {
	print_usage()
}

type CommandOptions struct {
	flags   *flag.FlagSet
	name    string
	summary string
	usage   string
}

var commands = []CommandOptions{}

func (c CommandOptions) print_usage() {
	l := log.New(os.Stderr, "", 0)

	l.Println(fmt.Sprintf("usage: git env %s\n", c.usage))
	l.Println(fmt.Sprintf("%s\n", c.summary))

	l.Println("global options:")
	flag.PrintDefaults()

	if c.flags.Name() != "" {
		l.Println("\noptions:")
		c.flags.PrintDefaults()
	}

	os.Exit(1)
}
