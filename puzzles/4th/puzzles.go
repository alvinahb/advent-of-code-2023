package dayFour

import (
	"math"
	"regexp"
	"strings"
)

// contains evaluates if a slice of strings contains a certain string
func contains(slice []string, element string) bool {
	for _, value := range slice {
		if value == element {
			return true
		}
	}
	return false
}

// getMatchingNumbers returns the number of matching numbers for a card
func getMatchingNumbers(winningNumbers string, numbersYouHave string) int {
	matchingNumbers := 0
	re := regexp.MustCompile("[0-9]+")

	for _, winningNumber := range re.FindAllString(winningNumbers, -1) {
		if contains(re.FindAllString(numbersYouHave, -1), winningNumber) {
			matchingNumbers += 1
		}
	}

	return matchingNumbers
}

// FirstPuzzle returns the result of the first puzzle
func FirstPuzzle() int {
	result := 0

	for _, card := range strings.Split(Input, "\n") {
		cardNumbers := strings.Split(card, ": ")[1]
		winningNumbers := strings.Split(cardNumbers, " | ")[0]
		numbersYouHave := strings.Split(cardNumbers, " | ")[1]

		matchingNumbers := getMatchingNumbers(winningNumbers, numbersYouHave)
		if matchingNumbers != 0 {
			result += int(math.Pow(2, float64(matchingNumbers-1)))
		}
	}

	return result
}

// SecondPuzzle returns the result of the second puzzle
func SecondPuzzle() int {
	result := 0
	cardCount := map[int]int{}

	for index, card := range strings.Split(Input, "\n") {
		if _, ok := cardCount[index]; !ok {
			cardCount[index] = 1
		}
		cardNumbers := strings.Split(card, ": ")[1]
		winningNumbers := strings.Split(cardNumbers, " | ")[0]
		numbersYouHave := strings.Split(cardNumbers, " | ")[1]

		matchingNumbers := getMatchingNumbers(winningNumbers, numbersYouHave)
		if matchingNumbers != 0 {
			// Add copies of future cards depending on the number of matching numbers
			for i := 1; i <= matchingNumbers; i++ {
				if _, ok := cardCount[index+i]; ok {
					cardCount[index+i] += cardCount[index]
				} else {
					cardCount[index+i] = cardCount[index] + 1
				}
			}
		}

		// Add cards to result
		result += cardCount[index]
	}

	return result
}
