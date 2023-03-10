package cards

func FavoriteCards() []int {
	return []int{2, 6, 9}
}

func GetItem(slice []int, index int) int {
	if outOfBounds(index, len(slice)) {
		return -1
	}

	return slice[index]
}

func SetItem(slice []int, index, value int) []int {
	if outOfBounds(index, len(slice)) {
		return append(slice, value)
	}

	slice[index] = value

	return slice
}

func PrependItems(slice []int, values ...int) []int {
	return append(values, slice...)
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	if outOfBounds(index, len(slice)) {
		return slice
	}

	return append(slice[:index], slice[index+1:]...)
}

func outOfBounds(index, sliceLength int) bool {
	return index < 0 || index > (sliceLength-1)
}
