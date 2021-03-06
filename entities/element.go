package entities

import "github.com/sidav/sidavgorandom/fibrandom"

const (
	ELEMENT_NONE = iota
	ELEMENT_FIRE
	ELEMENT_ICE
	ELEMENT_EARTH
	ELEMENT_AIR
	ELEMENT_MAGMA
	ELEMENT_STEAM
	ELEMENT_CHAOS
	ELEMENT_VOID

	frequencyOften = 7
	frequencyRare = 6
	frequencyEpic = 4
	frequencyLegendary = 1
)

type Element struct {
	Code int
}

// returns 1 if e2 is susceptible to e, 0 if unattuned, -1 if elements are attuned (i.e. reverse of susceptible)
func (e *Element) GetEffectivenessAgainstElement(e2 *Element) int {
	if e2.Code == ELEMENT_NONE {
		return 1 // everything is effective against no element
	}
	if e.Code == e2.Code {
		return -1
	}

	for _, e2code := range elementsTable[e.Code].susceptibleToDamageFrom {
		if e2.Code == e2code {
			return 1
		}
	}
	for _, e2code := range elementsTable[e.Code].isBoth {
		if e2.Code == e2code {
			return -1
		}
	}

	return 0
}

func (e *Element) GetName() string {
	return elementsTable[e.Code].name
}

func (e *Element) GetColorTags() []string {
	return elementsTable[e.Code].colorTags
}

func GetRandomElementCode(rnd *fibrandom.FibRandom) int {
	return rnd.SelectRandomIndexFromWeighted(len(elementsTable), func(x int) int { return elementsTable[x].frequency })
}

func GetRandomRareElementCode(rnd *fibrandom.FibRandom) int {
	return rnd.SelectRandomIndexFromWeighted(len(elementsTable),
		func(x int) int {
			if elementsTable[x].frequency == frequencyEpic || elementsTable[x].frequency == frequencyLegendary {
				return 1
			}
			return 0
		},
	)
}

type elementData struct {
	frequency int

	name                  string

	susceptibleToDamageFrom []int
	isBoth                  []int
	colorTags               []string
}

var elementsTable = []elementData{
	ELEMENT_NONE: {
		frequency:               frequencyOften,
		susceptibleToDamageFrom: []int{},
		isBoth:                  []int{},
	},
	ELEMENT_FIRE: {
		frequency:               frequencyRare,
		susceptibleToDamageFrom: []int{ELEMENT_ICE, ELEMENT_VOID},
		isBoth:                  []int{},
		colorTags:               []string{"RED"},
		name:                    "Flaming",
	},
	ELEMENT_ICE: {
		frequency:               frequencyRare,
		susceptibleToDamageFrom: []int{ELEMENT_FIRE, ELEMENT_VOID},
		isBoth:                  []int{},
		colorTags:               []string{"DARKCYAN"},
		name:                    "Ice",
	},
	ELEMENT_AIR: {
		frequency:               frequencyRare,
		susceptibleToDamageFrom: []int{ELEMENT_EARTH, ELEMENT_VOID},
		isBoth:                  []int{},
		colorTags:               []string{"YELLOW"},
		name:                    "Storm",
	},
	ELEMENT_EARTH: {
		frequency:               frequencyRare,
		susceptibleToDamageFrom: []int{ELEMENT_AIR, ELEMENT_VOID},
		isBoth:                  []int{},
		colorTags:               []string{"DARKGRAY"},
		name:                    "Stone",
	},
	ELEMENT_MAGMA: {
		frequency:               frequencyEpic,
		susceptibleToDamageFrom: []int{ELEMENT_ICE, ELEMENT_AIR, ELEMENT_STEAM, ELEMENT_VOID},
		isBoth:                  []int{ELEMENT_FIRE, ELEMENT_EARTH},
		colorTags:               []string{"RED", "DARKGRAY"},
		name:                    "Magmatic",
	},
	ELEMENT_STEAM: {
		frequency:               frequencyEpic,
		susceptibleToDamageFrom: []int{ELEMENT_FIRE, ELEMENT_EARTH, ELEMENT_MAGMA, ELEMENT_VOID},
		isBoth:                  []int{ELEMENT_AIR, ELEMENT_ICE},
		colorTags:               []string{"DARKCYAN", "YELLOW"},
		name:                    "Steaming",
	},
	ELEMENT_CHAOS: {
		frequency:               frequencyLegendary,
		susceptibleToDamageFrom: []int{ELEMENT_NONE, ELEMENT_VOID},
		isBoth:                  []int{ELEMENT_EARTH, ELEMENT_FIRE, ELEMENT_ICE, ELEMENT_AIR, ELEMENT_STEAM, ELEMENT_MAGMA},
		colorTags:               []string{"RED", "DARKCYAN", "YELLOW", "DARKGRAY"},
		name:                    "Chaotic",
	},
	ELEMENT_VOID: {
		frequency:               frequencyLegendary,
		susceptibleToDamageFrom: []int{ELEMENT_NONE, ELEMENT_CHAOS},
		isBoth:                  []int{ELEMENT_EARTH, ELEMENT_FIRE, ELEMENT_ICE, ELEMENT_AIR, ELEMENT_STEAM, ELEMENT_MAGMA},
		colorTags:               []string{"DARKGRAY", "DARKCYAN", "DARKMAGENTA"},
		name:                    "Void",
	},
}
