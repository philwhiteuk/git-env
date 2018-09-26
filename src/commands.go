package main

import (
	"bytes"
	"flag"
	"fmt"
)

type Command interface {
	Usage() bytes.Buffer
}

type CommandOptions struct {
	flags   *flag.FlagSet
	name    string
	summary string
	usage   string
}

func (c CommandOptions) Usage() bytes.Buffer {
	var b bytes.Buffer

	b.WriteString(fmt.Sprintf("usage: git env %s\n\n", c.usage))
	b.WriteString(fmt.Sprintf("%s\n\n", c.summary))

	if c.flags != nil {
		b.WriteString("options:\n")
		defaults := func(f *flag.Flag) {
			b.WriteString(fmt.Sprintf("   -%-11s %-14s\n", f.Name, f.Usage))
		}
		c.flags.VisitAll(defaults)
	}

	return b
}
