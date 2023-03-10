package lasagna

const (
	DefaultPrepTime         int     = 2
	Noodles                 string  = "noodles"
	NoodlesQuantityPerLayer int     = 50
	Sauce                   string  = "sauce"
	SauceQuantityPerLayer   float64 = 0.2
)

func PreparationTime(layers []string, avgPrepTimePerLayer int) int {
	if avgPrepTimePerLayer == 0 {
		avgPrepTimePerLayer = DefaultPrepTime
	}

	return len(layers) * avgPrepTimePerLayer
}

func Quantities(layers []string) (int, float64) {

	var noodles, sauce int

	for _, layer := range layers {
		switch layer {
		case Noodles:
			noodles++
		case Sauce:
			sauce++
		}
	}

	return noodles * NoodlesQuantityPerLayer, float64(sauce) * SauceQuantityPerLayer
}

func AddSecretIngredient(friendsList, myList []string) {
	secretIngredient := friendsList[len(friendsList)-1]

	myList[len(myList)-1] = secretIngredient
}

func ScaleRecipe(amountsNeeded []float64, numberOfPortions int) []float64 {
	var scaledRecipe []float64

	for _, amountNeeded := range amountsNeeded {
		scaledRecipe = append(scaledRecipe, (amountNeeded/2.0)*float64(numberOfPortions))
	}

	return scaledRecipe
}
