package main

const (
	DRF_BONFIRE = iota
	DRF_ALTAR
	DRF_TINKER
	DRF_BOSS
)

type dungeonRoomFeature struct {
	featureCode int
}

func getDungeonRoomFeatureNameByCode(x int) string {
	switch x {
	case DRF_BONFIRE: return "bonfire"
	case DRF_ALTAR: return "altar"
	case DRF_TINKER: return "tinkering table"
	case DRF_BOSS: return "Hydra Overlord remains"
	}
	panic("Y U NO IMPLEMENT")
}
