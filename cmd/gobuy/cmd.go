package main

import (
	"flag"
)

type Command struct {
	Run  func(cmd *Command, args ...string)
	Flag flag.FlagSet

	Name  string
	Usage string

	Summary string
	Help    string
}

func (c *Command) Exec(args []string) {
	c.Flag.Usage = func() {
		// helpFunc(c, c.Name)
	}
	c.Flag.Parse(args)
	c.Run(c, c.Flag.Args()...)
}
