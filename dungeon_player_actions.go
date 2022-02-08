package main

func (d *dungeon) selectPlayerRoomAction() {
	room := d.getPlayerRoom()
	var actions []string
	var allowed []bool
	actions = append(actions, "Nothing")
	allowed = append(allowed, true)
	actions = append(actions, "Pick up treasure")
	allowed = append(allowed, len(room.treasure)>0)
	actions = append(actions, "Rest")
	allowed = append(allowed, false)
	io.showSelectWindowWithDisableableOptions(
		"Select an action:", actions, func(x int) bool{return allowed[x]}, true)
}
