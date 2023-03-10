package purchase

var vehiclesNeedingLicense = map[string]bool{
	"car":   true,
	"truck": true,
}

func NeedsLicense(kind string) bool {
	return vehiclesNeedingLicense[kind]
}

func ChooseVehicle(option1, option2 string) string {
	betterChoice := option2

	if option1 < option2 {
		betterChoice = option1
	}

	return betterChoice + " is clearly the better choice."
}

func CalculateResellPrice(originalPrice, age float64) float64 {
	discountPercentage := 0.0

	switch {
	case age < 3:
		discountPercentage = 0.2
	case age < 10:
		discountPercentage = 0.3
	default:
		discountPercentage = 0.5
	}

	return calcuatePrice(originalPrice, discountPercentage)
}

func calcuatePrice(price, discountPercentage float64) float64 {
	return price * (1 - discountPercentage)
}
