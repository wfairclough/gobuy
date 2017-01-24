package main

import (
	"encoding/json"
	"fmt"

	"github.com/wfairclough/gobuy"
)

func main() {
	fmt.Println("Go Buy SDK command")

	client := gobuy.Client("example.myshopify.com", "Example App", "a8a5cd65ad764ac64d", 3)
	shop, err := client.GetShop()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	b, err := json.MarshalIndent(shop, "", "   ")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("%s\n\n", string(b))

	products, err := client.GetProducts(1)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	b, err = json.MarshalIndent(products, "", "   ")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("%s", string(b))
}
