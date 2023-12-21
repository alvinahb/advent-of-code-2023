package dayThree

import (
	"log"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

// Split input by lines
var lines = strings.Split(Input, "\n")

// isSymbol evaluates if a character is a symbol or not
func isSymbol(character rune) bool {
	additionalSymbols := []rune{'#', '*', '%', '@', '&', '-', '/'}

	if unicode.IsSymbol(character) || slices.Contains(additionalSymbols, character) {
		return true
	}

	return false
}

// isNextToSymbol returns if a number with given coordinates is adjacent to a symbol
func isNextToSymbol(coordinates [][]int) bool {
	for index, coordinate := range coordinates {
		x := coordinate[0]
		y := coordinate[1]

		var adjacentCoordinates = [][]int{}

		if index == 0 { // First element
			adjacentCoordinates = [][]int{
				{x, y - 1},     // Up
				{x - 1, y - 1}, // Up-left
				{x - 1, y},     // Left
				{x - 1, y + 1}, // Down-left
				{x, y + 1},     // Down
			}

		} else if index == len(coordinates)-1 { // Last element
			adjacentCoordinates = [][]int{
				{x, y - 1},     // Up
				{x + 1, y - 1}, // Up-right
				{x + 1, y},     // Right
				{x + 1, y + 1}, // Down-right
				{x, y + 1},     // Down
			}

		} else { // Intermediate element
			adjacentCoordinates = [][]int{
				{x, y - 1}, // Up
				{x, y + 1}, // Down
			}
		}

		for _, adjacentCoordinate := range adjacentCoordinates {
			adjX := adjacentCoordinate[0]
			adjY := adjacentCoordinate[1]
			// Check if coordinates are valid and if the character is a symbol
			if adjX >= 0 && adjY >= 0 && adjX < len(lines[0]) && adjY < len(lines) && isSymbol(rune(lines[adjY][adjX])) {
				return true
			}
		}
	}

	return false
}

// FirstPuzzle returns the result of the first puzzle
func FirstPuzzle() int {
	result := 0

	for y, line := range lines {
		// Retrieve numbers in line
		re := regexp.MustCompile("[0-9]+")
		numbersInString := re.FindAllString(line, -1)

		for _, numberString := range numbersInString {
			// Get number index
			index := strings.Index(line, numberString)
			numberCoordinates := [][]int{}

			// Get number coordinates
			for x := 0; x < len(numberString); x++ {
				numberCoordinates = append(numberCoordinates, []int{index + x, y})
			}

			// Evaluate if number is next to a symbol
			if isNextToSymbol(numberCoordinates) {
				if number, err := strconv.Atoi(numberString); err == nil {
					result += number
				} else {
					log.Fatal(err)
				}
			}
		}
	}

	return result
}
