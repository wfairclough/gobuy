package main

import (
	"encoding/json"
	"fmt"
	"log"
)

var productCmd = &Command{
	Name:    "product",
	Usage:   "",
	Summary: "Actions related to your Shopify products",
	Help:    `down extended help here...`,
	Run:     productRun,
}

func productRun(cmd *Command, args ...string) {

	client, err := clientFromConfig()
	if err != nil {
		log.Fatal(err)
	}

	products, err := client.GetProducts(1)
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
