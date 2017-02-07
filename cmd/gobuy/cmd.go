package main

import (
	"flag"
)

// Command Flag variables
var (
	page   *int
	limit  *int
	handle *string
)

type Command struct {
	Run  func(cmd *Command, args ...string)
	Flag flag.FlagSet

	Name  string
	Usage func(cmd *Command)

	Summary string
	Help    string

	SetupFlags func(fs *flag.FlagSet)
}

func (c *Command) Exec(args []string) {
	if c.SetupFlags != nil {
		c.SetupFlags(&c.Flag)
	}
	c.Flag.Usage = func() {
		c.Usage(c)
	}
	c.Flag.Parse(args)
	c.Run(c, c.Flag.Args()...)
}
