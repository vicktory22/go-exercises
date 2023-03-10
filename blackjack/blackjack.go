package blackjack

var cardValueMap = map[string]int{
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"ten":   10,
	"jack":  10,
	"queen": 10,
	"king":  10,
	"ace":   11,
}

func ParseCard(card string) int {
	return cardValueMap[card]
}

func FirstTurn(card1, card2, dealerCard string) string {
	playerValue := ParseCard(card1) + ParseCard(card2)
	dealerValue := ParseCard(dealerCard)

	switch {
	case isPairOfAces(playerValue):
		return "P"
	case isBlackJack(playerValue):
		return handleBlackJack(dealerValue)
	case inRange(17, 20, playerValue):
		return "S"
	case inRange(12, 16, playerValue):
		return handle12To16(dealerValue)
	default:
		return "H"
	}
}

func isPairOfAces(playerValue int) bool {
	return playerValue == 22
}

func isBlackJack(playerValue int) bool {
	return playerValue == 21
}

func handleBlackJack(dealerValue int) string {
	if inRange(10, 11, dealerValue) {
		return "S"
	}

	return "W"
}

func handle12To16(dealerValue int) string {
	if inRange(7, 11, dealerValue) {
		return "H"
	}

	return "S"
}

func inRange(min, max, value int) bool {
	return value >= min && value <= max
}
