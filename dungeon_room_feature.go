package main

const (
	DRF_BONFIRE = iota
	DRF_ALTAR
)

type dungeonRoomFeature struct {
	featureType int
}

func getDungeonRoomFeatureNameByCode(x int) string {
	switch x {
	case DRF_BONFIRE: return "bonfire"
	case DRF_ALTAR: return "altar"
	}
	panic("Y U NO IMPLEMENT")
}
