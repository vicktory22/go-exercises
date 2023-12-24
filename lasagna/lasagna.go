package lasagna

const OvenTime = 50
const TimePerLayer = 2

func RemainingOvenTime(actualMinutesInOven int) int {
	if actualMinutesInOven > OvenTime {
		panic("actualMinutesInOven cannot be greater than OvenTime")
	}

	return OvenTime - actualMinutesInOven
}

func PrepTime(numberOfLayers int) int {
	if numberOfLayers < 1 {
		panic("numberOfLayers cannot be less than 1")
	}

	return numberOfLayers * TimePerLayer
}

func ElapsedTime(numberOfLayers int, actualMinutesInOven int) int {
	return PrepTime(numberOfLayers) + actualMinutesInOven
}
