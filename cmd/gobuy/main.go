package main

import (
	"fmt"

	"github.com/wfairclough/gobuy"
)

func main() {
	fmt.Println("Go Buy SDK command")

	client := gobuy.Client("pibooths.myshopify.com", "pibooths", "apikey", 8)
	_ = client
}
