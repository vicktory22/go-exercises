package gross

var unitsMap = map[string]int{
	"quarter_of_a_dozen": 3,
	"half_of_a_dozen":    6,
	"dozen":              12,
	"small_gross":        120,
	"gross":              144,
	"great_gross":        1728,
}

func Units() map[string]int {
	return unitsMap
}

func NewBill() map[string]int {
	return map[string]int{}
}

func AddItem(bill, units map[string]int, item, unit string) bool {
	amount, exists := units[unit]

	if !exists {
		return false
	}

	bill[item] += amount

	return true
}

func RemoveItem(bill, units map[string]int, item, unit string) bool {
	billAmount, billExists := bill[item]
	unitAmount, unitExists := units[unit]

	if !billExists || !unitExists || billAmount < unitAmount {
		return false
	}

	bill[item] -= unitAmount

	if bill[item] == 0 {
		delete(bill, item)
	}

	return true
}

func GetItem(bill map[string]int, item string) (int, bool) {
	amount, exists := bill[item]

	return amount, exists
}
