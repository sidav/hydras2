package main

import "fmt"

var r cliRenderer

func main() {
	fmt.Printf("Nothing to see yet.\n")
	r.init()
	runGame()
	r.close()
	fmt.Printf("Works yet, though.\n")
}
