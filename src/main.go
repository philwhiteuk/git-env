package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
)

func main() {
	var command Command

	var help bool
	flag.BoolVar(&help, "h", false, "prints usage information")
	flag.Parse()

	r := ""
	if flag.NArg() > 0 {
		r = flag.Arg(0)
	}

	for _, c := range commands {
		if c.name == r {
			command = c
		}
	}

	if command != nil {
		command.print_usage()
		return
	} else {
		print_usage()
		if r != "" {
			fmt.Println(fmt.Sprintf("%s is not a recognized command", r))
		}
		os.Exit(1)
	}

}

func print_usage() {
	var usage bytes.Buffer

	usage.WriteString("usage: git env [<options>] <command> [<args>]\n\n")
	usage.WriteString("store and encrypt environment variables natively in git\n\n")

	for _, c := range commands {
		command_usage := fmt.Sprintf("   %-11s %-14s\n", c.name, c.summary)
		usage.WriteString(command_usage)
	}

	fmt.Println(usage.String())

	fmt.Println("global options:")
	flag.PrintDefaults()
}
