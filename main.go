package main

import "fmt"

var r cliRenderer

func main() {
	fmt.Printf("Nothing to see yet.\n")
	r.init()
	dung := &dungeon{}
	dung.initAndGenerate("patterns/explore_or_fight.ptn")
	r.renderDungeon(dung)
	r.close()
	fmt.Printf("Works yet, though.\n")
}
