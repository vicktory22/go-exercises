package grains

import (
	"errors"
	"math"
)

func Square(number int) (uint64, error) {
	if number < 1 || number > 64 {
		return 0, errors.New("Invalid input")
	}

	boardIndex := float64(number - 1)

	return uint64(math.Pow(2, boardIndex)), nil
}

func Total() uint64 {
	half64 := uint64(math.Pow(2, 63))

	return (half64 - 1) + half64
}
