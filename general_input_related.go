package main

func readKeyToVector(key string) (int, int) {
	switch key {
	case "UP", "w": return 0, -1
	case "DOWN", "s": return 0, 1
	case "LEFT", "a": return -1, 0
	case "RIGHT", "d": return 1, 0
	}
	return 0, 0
}
