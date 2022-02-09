package main

import "fmt"

func (d *dungeon) selectPlayerRoomAction() {
	room := d.getPlayerRoom()
	var actions []string
	var allowed []bool
	var actionFuncs []func()

	actions = append(actions, "Pick up treasure")
	allowed = append(allowed, len(room.treasure)>0)
	actionFuncs = append(actionFuncs, d.pickUpFromPlayerRoom)

	actions = append(actions, "Rest")
	allowed = append(allowed, false)
	actionFuncs = append(actionFuncs, func(){})

	actions = append(actions, "Nothing")
	allowed = append(allowed, true)
	actionFuncs = append(actionFuncs, func(){})

	chosenActionNum := io.showSelectWindowWithDisableableOptions(
		"Select an action:", actions, func(x int) bool{return allowed[x]}, false)
	if chosenActionNum != -1 {
		actionFuncs[chosenActionNum]()
	}
}

func (d *dungeon) pickUpFromPlayerRoom() {
	room := d.getPlayerRoom()
	for len(room.treasure) > 0 {
		lines := []string{}
		for _, i := range room.treasure {
			if i.IsConsumable() {
				lines = append(lines, fmt.Sprintf("%dx %s", i.AsConsumable.Amount, i.GetName()))
			} else {
				lines = append(lines, fmt.Sprintf(i.GetName()))
			}
		}
		picked := io.showSelectWindow("Pick up:", lines)
		if picked == -1 {
			break
		}
		d.plr.acquireItem(room.treasure[picked])
		room.treasure = append(room.treasure[:picked], room.treasure[picked+1:]...)
	}
}
