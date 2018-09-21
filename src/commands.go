package main

import (
	"flag"
	"fmt"
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
	fmt.Println(fmt.Sprintf("usage: git env %s\n", c.usage))
	fmt.Println(fmt.Sprintf("%s\n", c.summary))

	fmt.Println("global options:")
	flag.PrintDefaults()

	if c.flags.Name() != "" {
		fmt.Println("\noptions:")
		c.flags.PrintDefaults()
	}

}
