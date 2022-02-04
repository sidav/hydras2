package main

import generator "cyclicdungeongenerator/generators"

type dungeon struct {
	name           string
	layout         generator.LayoutInterface
	rooms          [][]*room
	totalStages    int
}
