package main

func readKeyToVector(key string) (int, int) {
	switch key {
	case "UP": return 0, -1
	case "DOWN": return 0, 1
	case "LEFT": return -1, 0
	case "RIGHT": return 1, 0
	}
	return 0, 0
}
