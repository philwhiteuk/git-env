package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
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
			l := log.New(os.Stderr, "", 0)
			l.Println(fmt.Sprintf("\n%s is not a recognized command", r))
		}
		os.Exit(1)
	}
}

func print_usage() {
	l := log.New(os.Stderr, "", 0)

	l.Println(fmt.Sprintf("usage: git env [<options>] <command> [<args>]\n"))
	l.Println(fmt.Sprintf("store and encrypt environment variables natively in git"))

	var b bytes.Buffer
	for _, c := range commands {
		b.WriteString(fmt.Sprintf("   %-11s %-14s\n", c.name, c.summary))
	}
	l.Println(b.String())

	l.Println("global options:")
	flag.PrintDefaults()
}
