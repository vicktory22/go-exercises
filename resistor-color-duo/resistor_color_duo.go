package resistorcolorduo

// import (
// 	"fmt"
// 	"strconv"
// )

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

func Value(colors []string) int {
	if len(colors) < 2 {
		panic("Need at least 2 colors")
	}

	firstColor := resistorMap[colors[0]]
	secondColor := resistorMap[colors[1]]

	return (firstColor * 10) + secondColor
}
