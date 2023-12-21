package dayOne

import (
	"log"
	"strconv"
	"strings"
)

// List of numbers in string
var numbers = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

// FirstPuzzle returns the result of the first puzzle
func FirstPuzzle() int {
	result := 0

	for _, inputLine := range strings.Split(Input, "\n") {
		firstDigitIndex := -1
		secondDigitIndex := -1

		for _, number := range numbers {
			// Get first occurence index of number if existing
			index := strings.Index(inputLine, number)
			if index != -1 && (firstDigitIndex == -1 || index < firstDigitIndex) {
				firstDigitIndex = index
			}

			// Get last occurence index of number if existing
			index = strings.LastIndex(inputLine, number)
			if index != -1 && (secondDigitIndex == -1 || index > secondDigitIndex) {
				secondDigitIndex = index
			}
		}

		// Calculate calibration value and add it to result
		if firstDigitIndex != -1 && secondDigitIndex != -1 {
			stringDigit := string(inputLine[firstDigitIndex]) + string(inputLine[secondDigitIndex])
			digit, err := strconv.Atoi(stringDigit)
			if err != nil {
				log.Fatal(err)
			}
			result += digit
		}
	}

	return result
}

// Maps numbers from letters to digit
var numbersInLetters = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

// SecondPuzzle returns the result of the second puzzle
func SecondPuzzle() int {
	result := 0

	for _, inputLine := range strings.Split(Input, "\n") {
		firstDigitIndex := -1
		secondDigitIndex := -1
		firstDigit := ""
		secondDigit := ""

		for _, number := range numbers {
			// Get first occurence index of number if existing
			index := strings.Index(inputLine, number)
			if index != -1 && (firstDigitIndex == -1 || index < firstDigitIndex) {
				firstDigitIndex = index
				firstDigit = string(inputLine[index])
			}

			// Get last occurence index of number if existing
			index = strings.LastIndex(inputLine, number)
			if index != -1 && (secondDigitIndex == -1 || index > secondDigitIndex) {
				secondDigitIndex = index
				secondDigit = string(inputLine[index])
			}
		}

		for key, value := range numbersInLetters {
			// Get first occurence index of number if existing
			index := strings.Index(inputLine, key)
			if index != -1 && (firstDigitIndex == -1 || index < firstDigitIndex) {
				firstDigitIndex = index
				firstDigit = value
			}

			// Get last occurence index of number if existing
			index = strings.LastIndex(inputLine, key)
			if index != -1 && (secondDigitIndex == -1 || index > secondDigitIndex) {
				secondDigitIndex = index
				secondDigit = value
			}
		}

		// Calculate calibration value and add it to result
		if firstDigit != "" && secondDigit != "" {
			stringDigit := firstDigit + secondDigit
			digit, err := strconv.Atoi(stringDigit)
			if err != nil {
				log.Fatal(err)
			}
			result += digit
		}
	}

	return result
}
