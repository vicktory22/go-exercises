package lasagna

const FortyMinutes = 40
const TwoMinutes = 2

const OvenTime = FortyMinutes
const PreparationTimePerLayer = TwoMinutes

func RemainingOvenTime(actualMinutesInOven int) int {
  return OvenTime - actualMinutesInOven
}

func PreparationTime(numberOfLayers int) int {
  return PreparationTimePerLayer * numberOfLayers
}

func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
  return PreparationTime(numberOfLayers) + actualMinutesInOven 
}
