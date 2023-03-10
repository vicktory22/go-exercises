package cars

const PRICE_PER_CAR = 10000
const DISCOUNT_GROUP_SIZE = 10
const DISCOUNT_RATE = 0.05

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * (successRate / 100.00)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	perHourRate := CalculateWorkingCarsPerHour(productionRate, successRate)

	return int(perHourRate) / 60
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	discountGroupCount := carsCount / DISCOUNT_GROUP_SIZE
	singleCarCount := carsCount % DISCOUNT_GROUP_SIZE

	discountGroupTotal := float64(discountGroupCount * PRICE_PER_CAR * DISCOUNT_GROUP_SIZE)

	discountGroupCost := discountGroupTotal - (discountGroupTotal * DISCOUNT_RATE)
	singleCartCost := singleCarCount * PRICE_PER_CAR

	return uint(discountGroupCost) + uint(singleCartCost)
}
