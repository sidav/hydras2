package main

import (
	"fmt"
	"hydras2/entities"
)

func readKeyToVector(key string) (int, int) {
	switch key {
	case "UP", "w":
		return 0, -1
	case "DOWN", "s":
		return 0, 1
	case "LEFT", "a":
		return -1, 0
	case "RIGHT", "d":
		return 1, 0
	}
	return 0, 0
}

func getAttackDescriptionString(weapon *entities.Item, target *enemy) string {
	attackStr := fmt.Sprintf("%d", target.heads)
	damageNum := weapon.AsWeapon.GetDamageOnHeads(target.heads)
	switch weapon.AsWeapon.WeaponTypeCode {
	case entities.WTYPE_DIVISOR:
		attackStr += fmt.Sprintf("/%d", weapon.AsWeapon.Damage)
	default:
		attackStr += fmt.Sprintf("-%d", damageNum)
	}
	remainingHeadsNum := target.heads - damageNum
	if remainingHeadsNum == 0 {
		attackStr += "->kill"
	} else {
		switch weapon.AsWeapon.WeaponElement.GetEffectivenessAgainstElement(target.element) {
		case -1: attackStr = "("+attackStr+fmt.Sprintf(")*2=%d", remainingHeadsNum*2)
		case 0: attackStr += fmt.Sprintf("+1=%d", remainingHeadsNum+1)
		case 1: attackStr += fmt.Sprintf("=%d", remainingHeadsNum)
		}
	}
	return attackStr
}
