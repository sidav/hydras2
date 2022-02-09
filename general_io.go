package main

import "hydras2/entities"

func selectAnItemFromList(title string, itms []*entities.Item) *entities.Item {
	var lines []string
	for _, i := range itms {
		lines = append(lines, i.GetName())
	}
	index := io.showSelectWindow(title, lines)
	if index == -1 {
		return nil
	}
	return itms[index]
}
