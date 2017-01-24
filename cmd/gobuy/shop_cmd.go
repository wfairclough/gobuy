package main

import (
	"encoding/json"
	"fmt"
	"log"
)

var shopCmd = &Command{
	Name:    "shop",
	Usage:   "",
	Summary: "Actions related to your Shopify Shop",
	Help:    `down extended help here...`,
	Run:     shopRun,
}

func shopRun(cmd *Command, args ...string) {

	client, err := clientFromConfig()
	if err != nil {
		log.Fatal(err)
	}

	shop, err := client.GetShop()
	if err != nil {
		log.Println(err.Error())
		return
	}
	b, err := json.MarshalIndent(shop, "", "   ")
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	fmt.Printf("%s\n\n", string(b))
}
