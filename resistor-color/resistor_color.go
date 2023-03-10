package resistorcolor

var resistorMap = map[string]int{
	"black":  0,
	"brown":  1,
	"red":    2,
	"orange": 3,
	"yellow": 4,
	"green":  5,
	"blue":   6,
	"violet": 7,
	"grey":   8,
	"white":  9,
}

func Colors() []string {
	keys := make([]string, 0, len(resistorMap))

	for key := range resistorMap {
		keys = append(keys, key)
	}

	return keys
}

func ColorCode(color string) int {
	colorCode, exists := resistorMap[color]

	if !exists {
		panic("Unknown color code given")
	}

	return colorCode
}
