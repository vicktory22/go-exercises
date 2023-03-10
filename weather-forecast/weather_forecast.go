// Package weather is a package that provides the weather forecast.
package weather

// CurrentCondition is the current weather condition.
var CurrentCondition string

// CurrentLocation is the current location of the user.
var CurrentLocation string

// Forecast takes in a city and condition and returns a formatted weather forecast.
func Forecast(city string, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
