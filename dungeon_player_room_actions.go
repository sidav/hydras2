package main

import (
	"fmt"
	"hydras2/entities"
)

func (d *dungeon) selectPlayerRoomAction() {
	room := d.getPlayerRoom()
	var actions []string
	var allowed []bool
	var actionFuncs []func()

	actions = append(actions, "Pick up treasure")
	allowed = append(allowed, len(room.treasure)>0)
	actionFuncs = append(actionFuncs, d.pickUpFromPlayerRoom)

	actions = append(actions, "Rest")
	allowed = append(allowed, room.hasFeature(DRF_BONFIRE))
	actionFuncs = append(actionFuncs, d.playerRest)

	actions = append(actions, "Pray")
	allowed = append(allowed, room.hasFeature(DRF_ALTAR))
	actionFuncs = append(actionFuncs, d.buyPlayerStatUpgrades)

	actions = append(actions, "Craft")
	allowed = append(allowed, room.hasFeature(DRF_TINKER))
	actionFuncs = append(actionFuncs, d.tinkerWithItems)

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

func (d *dungeon) playerRest() {
	lines := []string{
		"Rest for some time?",
		"It will restore your health, but the",
		"dungeon will become dark again, and cleared",
		"rooms may be repopulated with hydras!",
	}
	picked := io.showYNSelect("REST", lines)
	if picked {
		d.plr.hitpoints = d.plr.getMaxHp()
		d.clearRoomsGeneratedState()
	}
}

func (d *dungeon) buyPlayerStatUpgrades() {
	upgradeCost := d.plr.level * 75 / 10

	if d.plr.souls < upgradeCost {
		io.showInfoWindow("NOTHING TO OFFER", []string{
			fmt.Sprintf("You need %d more essence to offer.", upgradeCost - d.plr.souls),
		})
		return
	}

	lines := []string{
		fmt.Sprintf("You have %d hydra essense.", d.plr.souls),
		fmt.Sprintf("Spend %d to upgrade vitality?", upgradeCost),
	}
	picked := io.showYNSelect("MAKE AN OFFERING", lines)
	if picked && d.plr.souls >= upgradeCost {
		d.plr.souls -= upgradeCost
		d.plr.vitality += 2
		d.plr.level += 1
	}
}

func (d *dungeon) tinkerWithItems() {
	if len(d.plr.getAllMaterials()) == 0 {
		io.showInfoWindow("NO MATERIALS", []string{"You have no materials to tinker with."})
		return
	}
	selectedMaterial := selectAnItemFromList("SELECT MATERIAL:", d.plr.getAllMaterials())
	if selectedMaterial == nil {
		return
	}

	var selectedItem *entities.Item
	// upgrading flask is a separate thing
	if selectedMaterial.AsMaterial.Code == entities.MATERIAL_ENCHANT_CONSUMABLE {
		itms := d.plr.getAllItemsOfTypeAndCode(entities.ITEM_CONSUMABLE, entities.CONSUMABLE_HEAL)
		if len(itms) > 1 {
			panic("What? More than one flask?!")
		}
		selectedItem = itms[0]
	} else { // weapon
		selectedItem = selectAnItemFromList("SELECT WEAPON TO APPLY "+selectedMaterial.GetName(), d.plr.getAllWeapons())
		if selectedItem == nil {
			return
		}
	}
	preTinkerWeaponName := selectedItem.GetName()
	applyMaterialToItem(selectedMaterial, selectedItem)
	io.showInfoWindow("SUCCESS", []string{
		"You made ",
		preTinkerWeaponName,
		"into",
		selectedItem.GetName()},
	)
	d.plr.removeItemFromInventory(selectedMaterial)
}
