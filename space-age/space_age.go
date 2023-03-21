package space

type Planet string

const earthSecondsPerYear = 31557600

var planetSeconds = map[Planet]float64{
	"Mercury": earthSecondsPerYear * 0.2408467,
	"Venus":   earthSecondsPerYear * 0.61519726,
	"Earth":   earthSecondsPerYear,
	"Mars":    earthSecondsPerYear * 1.8808158,
	"Jupiter": earthSecondsPerYear * 11.862615,
	"Saturn":  earthSecondsPerYear * 29.447498,
	"Uranus":  earthSecondsPerYear * 84.016846,
	"Neptune": earthSecondsPerYear * 164.79132,
}

func Age(seconds float64, planet Planet) float64 {
	secondsPerYear, ok := planetSeconds[planet]

	if !ok {
		return -1.0
	}

	return seconds / secondsPerYear
}
