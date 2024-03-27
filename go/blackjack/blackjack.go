package blackjack

var stringNumberMap = map[string]int{
	"ace":   11,
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
	"joker": 0,
}

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	return stringNumberMap[card]
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	card1Val := ParseCard(card1)
	card2Val := ParseCard(card2)
	cardValsSum := card1Val + card2Val
	dealerCardVal := ParseCard(dealerCard)
	switch {
	case card1Val == ParseCard("ace") && card2Val == ParseCard("ace"):
		return "P"
	case cardValsSum >= 21 && dealerCardVal < 10:
		return "W"
	case cardValsSum >= 21 && dealerCardVal >= 10:
		return "S"
	case cardValsSum >= 17 && cardValsSum <= 20:
		return "S"
	case cardValsSum >= 12 && cardValsSum <= 16 && dealerCardVal < 7:
		return "S"
	}
	return "H"
}
