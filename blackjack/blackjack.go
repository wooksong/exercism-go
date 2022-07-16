package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	var retVal int

	switch card {
		case "ace":
			retVal = 11
		case "two":
			retVal = 2
		case "three":
			retVal = 3
		case "four":
			retVal = 4
		case "five":
			retVal = 5
		case "six":
			retVal = 6
		case "seven":
			retVal = 7
		case "eight":
			retVal = 8
		case "nine":
			retVal = 9
		case "ten":
			fallthrough
		case "jack":
			fallthrough
		case "queen":
			fallthrough
		case "king":
			retVal = 10
		default:
			retVal = 0
	}

	return retVal
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	var action string

	numCard1 := ParseCard(card1)
	numCard2 := ParseCard(card2)
	numDealerCard := ParseCard(dealerCard)
	sumCards := numCard1 + numCard2

	switch {
		case numCard1 == 11 && numCard2 == 11:
			action = "P"
		case sumCards == 21:
			if numDealerCard >= 10 {
				action = "S"
			} else {
				action = "W"
			}
		case sumCards >= 17 && sumCards <= 20:
			action = "S"
		case sumCards >= 12 && sumCards <= 16:
			if numDealerCard >= 7 {
				action = "H"
			} else {
				action = "S"
			}
		case sumCards <= 11:
			action = "H"
	}

	return action
}
