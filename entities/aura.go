package entities

import "github.com/sidav/sidavgorandom/fibrandom"

// Aura is like "brand for enemy"
type Aura struct {
	Code int
}

func (a *Aura) GetName() string {
	return aurasTable[a.Code].name
}

const (
	AURA_SELF_GROWING = iota
	AURA_OTHERS_GROWING
	AURA_FAST
	AURA_SUMMONING
	AURA_VAMPIRIC
)

type auraData struct {
	name      string
	frequency int
}

var aurasTable = map[int]*auraData{
	AURA_SELF_GROWING: {
		name: "fast-growing",
	},
	AURA_OTHERS_GROWING: {
		name: "reinforcing",
	},
	AURA_FAST: {
		name: "fast",
	},
	AURA_SUMMONING: {
		name: "summoner",
	},
	AURA_VAMPIRIC: {
		name: "vampiric",
	},
}

func GenerateRandomAura(rnd *fibrandom.FibRandom) *Aura {
	return &Aura{
		Code: rnd.SelectRandomIndexFromWeighted(len(aurasTable), func(x int) int {return 1}),
	}
}
