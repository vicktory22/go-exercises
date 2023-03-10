package resistorcolortrio

import (
	"fmt"
)

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

var magnitudes = [10]string{
	" ",
	"0 ",
	" kilo",
	" kilo",
	"0 kilo",
	" mega",
	" mega",
	"0 mega",
	" giga",
	" giga",
}

func Label(colors []string) string {
	if len(colors) < 3 {
		panic("Need at least 3 colors")
	}

	first := resistorMap[colors[0]]
	second := resistorMap[colors[1]]
	third := resistorMap[colors[2]]

	ohms := (first * 10) + second

	if second == 0 {
		ohms = ohms / 10
	}

	return fmt.Sprintf("%d%sohms", ohms, magnitudes[third])

}
