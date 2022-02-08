package entities

import "github.com/sidav/sidavgorandom/fibrandom"

const (
	ELEMENT_NONE = iota
	ELEMENT_FIRE
	ELEMENT_ICE
	ELEMENT_EARTH
	ELEMENT_AIR
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
	for _, e2code := range elementsTable[e.Code].attunedWith {
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

func GetWeightedRandomElementCode(rnd *fibrandom.FibRandom) int {
	return rnd.SelectRandomIndexFromWeighted(len(elementsTable), func(x int) int { return elementsTable[x].frequency })
}

type elementData struct {
	frequency int
	name      string

	susceptibleToDamageFrom []int
	attunedWith             []int
	colorTags               []string
}

var elementsTable = []elementData{
	ELEMENT_NONE: {
		frequency:               8,
		susceptibleToDamageFrom: []int{},
		attunedWith:             []int{},
	},
	ELEMENT_FIRE: {
		frequency:               1,
		susceptibleToDamageFrom: []int{ELEMENT_ICE},
		attunedWith:             []int{},
		colorTags:               []string{"RED"},
		name:                    "Flaming",
	},
	ELEMENT_ICE: {
		frequency:               1,
		susceptibleToDamageFrom: []int{ELEMENT_FIRE},
		attunedWith:             []int{},
		colorTags:               []string{"BLUE"},
		name:                    "Ice",
	},
	ELEMENT_AIR: {
		frequency:               1,
		susceptibleToDamageFrom: []int{ELEMENT_EARTH},
		attunedWith:             []int{},
		colorTags:               []string{"CYAN"},
		name:                    "Storm",
	},
	ELEMENT_EARTH: {
		frequency:               1,
		susceptibleToDamageFrom: []int{ELEMENT_AIR},
		attunedWith:             []int{},
		colorTags:               []string{"DARKGRAY"},
		name:                    "Stone",
	},
}
