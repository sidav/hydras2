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

func GetWeightedRandomElementCode(rnd *fibrandom.FibRandom) int {
	return rnd.SelectRandomIndexFromWeighted(len(elementsTable), func(x int) int { return elementsTable[x].frequencyUsual })
}

type elementData struct {
	frequencyUsual        int
	frequencyIfPreferRare int
	frequencyIfPreferEpic int
	name                  string

	susceptibleToDamageFrom []int
	isBoth                  []int
	colorTags               []string
}

var elementsTable = []elementData{
	ELEMENT_NONE: {
		frequencyUsual:          8,
		frequencyIfPreferRare:   4,
		frequencyIfPreferEpic:   0,
		susceptibleToDamageFrom: []int{},
		isBoth:                  []int{},
	},
	ELEMENT_FIRE: {
		frequencyUsual:          2,
		frequencyIfPreferRare:   4,
		frequencyIfPreferEpic:   1,
		susceptibleToDamageFrom: []int{ELEMENT_ICE},
		isBoth:                  []int{},
		colorTags:               []string{"RED"},
		name:                    "Flaming",
	},
	ELEMENT_ICE: {
		frequencyUsual:          2,
		frequencyIfPreferRare:   4,
		frequencyIfPreferEpic:   1,
		susceptibleToDamageFrom: []int{ELEMENT_FIRE},
		isBoth:                  []int{},
		colorTags:               []string{"BLUE"},
		name:                    "Ice",
	},
	ELEMENT_AIR: {
		frequencyUsual:          2,
		frequencyIfPreferRare:   4,
		frequencyIfPreferEpic:   1,
		susceptibleToDamageFrom: []int{ELEMENT_EARTH},
		isBoth:                  []int{},
		colorTags:               []string{"YELLOW"},
		name:                    "Storm",
	},
	ELEMENT_EARTH: {
		frequencyUsual:          2,
		frequencyIfPreferRare:   4,
		frequencyIfPreferEpic:   1,
		susceptibleToDamageFrom: []int{ELEMENT_AIR},
		isBoth:                  []int{},
		colorTags:               []string{"DARKGRAY"},
		name:                    "Stone",
	},
	ELEMENT_MAGMA: {
		frequencyUsual:          1,
		frequencyIfPreferRare:   4,
		frequencyIfPreferEpic:   1,
		susceptibleToDamageFrom: []int{ELEMENT_ICE, ELEMENT_AIR, ELEMENT_STEAM},
		isBoth:                  []int{ELEMENT_FIRE, ELEMENT_EARTH},
		colorTags:               []string{"RED", "DARKGRAY"},
		name:                    "Magmatic",
	},
	ELEMENT_STEAM: {
		frequencyUsual:          1,
		frequencyIfPreferRare:   4,
		frequencyIfPreferEpic:   2,
		susceptibleToDamageFrom: []int{ELEMENT_FIRE, ELEMENT_EARTH, ELEMENT_MAGMA},
		isBoth:                  []int{ELEMENT_AIR, ELEMENT_ICE},
		colorTags:               []string{"BLUE", "YELLOW"},
		name:                    "Steaming",
	},
}
