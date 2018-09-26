package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
)

var commands = []CommandOptions{}

func main() {
	var command Command

	var help bool
	flag.BoolVar(&help, "-help", false, "prints usage information")
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

	l := log.New(os.Stderr, "", 0)
	if command != nil {
		b := command.Usage()
		l.Println(b.String())
	} else {
		print_usage()
		if r != "" {
			e := errors.New(fmt.Sprintf("%s not a valid command", r))
			l.Println(fmt.Sprintf("\n%s", e))
		}
	}
	os.Exit(1)
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

	flag.PrintDefaults()
}
