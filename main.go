package main

import (
	"fmt"
	"./pkg/hostlist-creator"
	"./pkg/create-hosts"
)

func main() {
	numNotebooks := 10 // Number of notebooks to generate

	hostlistcreator.GenerateHostList(numNotebooks)
	fmt.Println("Host list generated successfully")
}