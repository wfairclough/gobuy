package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var collectionCmd = &Command{
	Name:    "collection",
	Usage:   collectionUsage,
	Summary: "Actions related to your Shopify collections",
	Help:    `down extended help here...`,
	Run:     collectionRun,
}

var collectionCommands = []*Command{
	collectionListCmd,
	collectionGetCmd,
}

func collectionRun(cmd *Command, args ...string) {

	if len(args) == 0 {
		cmd.Flag.Usage()
		return
	}

	var subcmd *Command
	name := args[0]
	for _, c := range collectionCommands {
		if strings.HasPrefix(c.Name, name) {
			subcmd = c
			break
		}
	}

	subcmd.Exec(args[1:])
}

func collectionUsage(cmd *Command) {
	fmt.Print(collectionUsagePrefix)
	cmd.Flag.PrintDefaults()
	usageTmpl.Execute(os.Stdout, collectionCommands)
}

var collectionUsagePrefix = `
The collection command supplies actions related to your Shopify Collections

Usage:
    gobuy [options] collection <subcommand> [subcommand options]

Options:
`

var collectionGetCmd = &Command{
	Name:    "get",
	Usage:   collectionGetUsage,
	Summary: "Get a specific collection",
	Help:    `down extended help here...`,
	Run:     collectionGetRun,
	SetupFlags: func(fs *flag.FlagSet) {
		handle = fs.String("handle", "", "The handle for the collection to fetch")
	},
}

func collectionGetRun(cmd *Command, args ...string) {
	// if len(args) == 0 {
	// 	cmd.Flag.Usage()
	// 	return
	// }

	client, err := clientFromConfig()
	if err != nil {
		log.Fatal(err)
	}

	collections, err := client.GetCollectionByHandle(*handle)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	b, err := json.MarshalIndent(collections, "", "   ")
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	fmt.Printf("%s\n\n", string(b))
}

func collectionGetUsage(cmd *Command) {
	fmt.Print(collectionGetUsagePrefix)
	cmd.Flag.PrintDefaults()
	// usageTmpl.Execute(os.Stdout, collectionCommands)
}

var collectionGetUsagePrefix = `
The collection get command supplies actions related to your Shopify Collections

Usage:
    gobuy [options] collection get [get options]

Options:
`

var collectionListCmd = &Command{
	Name:    "list",
	Usage:   collectionListUsage,
	Summary: "List collections by page",
	Help:    `down extended help here...`,
	Run:     collectionListRun,
	SetupFlags: func(fs *flag.FlagSet) {
		page = fs.Int("page", 1, "The page to list of collections")
		limit = fs.Int("limit", 10, "The limit of amount of collections to list")
	},
}

func collectionListRun(cmd *Command, args ...string) {
	// if len(args) == 0 || args[0] == "-h" {
	// 	cmd.Flag.Usage()
	// 	return
	// }

	client, err := clientFromConfig()
	if err != nil {
		log.Fatal(err)
	}

	collections, err := client.GetCollections(*page, *limit)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	b, err := json.MarshalIndent(collections, "", "   ")
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	fmt.Printf("%s\n\n", string(b))
}

func collectionListUsage(cmd *Command) {
	fmt.Print(collectionListUsagePrefix)
	cmd.Flag.PrintDefaults()
	// usageTmpl.Execute(os.Stdout, collectionCommands)
}

var collectionListUsagePrefix = `
The collection list command supplies actions related to listing Shopify Collections

Usage:
    gobuy [options] collection list [list options]

Options:
`
