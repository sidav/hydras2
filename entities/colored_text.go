package entities

const (
	COLOR_TAG_LENGTH = 5
	TAG_SYMBOL = "@"
)

var ColorTagsTable = map[string]string{
	"RED":      "@red@",
	"BLUE":     "@blu@",
	"CYAN":     "@cyn@",
	"DARKGRAY": "@dgr@",
	"RESET":    "@nil@",
}

func GetColorTagNameInStringAtPosition(s string, pos int) string {
	if len(s)-pos < COLOR_TAG_LENGTH {
		return ""
	}
	potentialTag := (s[pos:])[:COLOR_TAG_LENGTH]
	for k, v := range ColorTagsTable {
		if len(v) != len(potentialTag) {
			panic("Tag length error")
		}
		if potentialTag == v {
			return k
		}
	}
	return ""
}

func IsStringColorTagged(s string) bool {
	for i := 0; i < len(s)-COLOR_TAG_LENGTH; i++ {
		if string(s[i]) == TAG_SYMBOL {
			return true
		}
	}
	return false
}

func MakeStringColorTagged(s string, tagsNames []string) string {
	switch len(tagsNames) {
	case 0:
		return s
	case 1:
		return ColorTagsTable[tagsNames[0]] + s
	}
	// maybe calculate this when rendering?.. Why consume memory?
	newStr := ""
	const step = 2
	for i := 0; i < len(s); i++ {
		if i % step == 0 {
			newStr += ColorTagsTable[tagsNames[(i/step)%len(tagsNames)]]
		}
		newStr += string(s[i])
	}
	return newStr
	panic("Y U NO IMPLEMENT")
}
