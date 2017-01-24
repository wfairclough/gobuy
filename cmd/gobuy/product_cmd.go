package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var productCmd = &Command{
	Name:    "product",
	Usage:   productUsage,
	Summary: "Actions related to your Shopify products",
	Help:    `down extended help here...`,
	Run:     productRun,
}

var productCommands = []*Command{
	productListCmd,
	productGetCmd,
}

func productRun(cmd *Command, args ...string) {

	if len(args) == 0 {
		cmd.Flag.Usage()
		return
	}

	var subcmd *Command
	name := args[0]
	for _, c := range productCommands {
		if strings.HasPrefix(c.Name, name) {
			subcmd = c
			break
		}
	}

	subcmd.Exec(args[1:])
}

func productUsage(cmd *Command) {
	fmt.Print(productUsagePrefix)
	cmd.Flag.PrintDefaults()
	usageTmpl.Execute(os.Stdout, productCommands)
}

var productUsagePrefix = `
The product command supplies actions related to your Shopify Products

Usage:
    gobuy [options] product <subcommand> [subcommand options]

Options:
`

var productGetCmd = &Command{
	Name:    "get",
	Usage:   productGetUsage,
	Summary: "Get a specific product",
	Help:    `down extended help here...`,
	Run:     productGetRun,
	SetupFlags: func(fs *flag.FlagSet) {
		fs.String("handle", "", "The handle for the product to fetch")
	},
}

func productGetRun(cmd *Command, args ...string) {
	if len(args) == 0 {
		cmd.Flag.Usage()
		return
	}

	log.Fatal("Unimplemented product get")
}

func productGetUsage(cmd *Command) {
	fmt.Print(productGetUsagePrefix)
	cmd.Flag.PrintDefaults()
	// usageTmpl.Execute(os.Stdout, productCommands)
}

var productGetUsagePrefix = `
The product get command supplies actions related to your Shopify Products

Usage:
    gobuy [options] product get [get options]

Options:
`

var page, limit *int
var productListCmd = &Command{
	Name:    "list",
	Usage:   productListUsage,
	Summary: "List products by page",
	Help:    `down extended help here...`,
	Run:     productListRun,
	SetupFlags: func(fs *flag.FlagSet) {
		page = fs.Int("page", 1, "The page to list of products")
		limit = fs.Int("limit", 10, "The limit of amount of products to list")
	},
}

func productListRun(cmd *Command, args ...string) {
	// if len(args) == 0 || args[0] == "-h" {
	// 	cmd.Flag.Usage()
	// 	return
	// }

	client, err := clientFromConfig()
	if err != nil {
		log.Fatal(err)
	}

	products, err := client.GetProducts(*page, *limit)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	b, err := json.MarshalIndent(products, "", "   ")
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	fmt.Printf("%s\n\n", string(b))
}

func productListUsage(cmd *Command) {
	fmt.Print(productListUsagePrefix)
	cmd.Flag.PrintDefaults()
	// usageTmpl.Execute(os.Stdout, productCommands)
}

var productListUsagePrefix = `
The product list command supplies actions related to listing Shopify Products

Usage:
    gobuy [options] product list [list options]

Options:
`
