package birdwatcher

func TotalBirdCount(birdsPerDay []int) int {
	total := 0

	for _, dailyCount := range birdsPerDay {
		total += dailyCount
	}

	return total
}

func BirdsInWeek(birdsPerDay []int, week int) int {
	daysInWeek := 7
	end := week * daysInWeek
	start := end - daysInWeek

	return TotalBirdCount(birdsPerDay[start:end])
}

func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i += 2 {
		birdsPerDay[i]++
	}

	return birdsPerDay
}
