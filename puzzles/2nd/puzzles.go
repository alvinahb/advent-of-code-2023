package dayTwo

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

// Maximum of cubes by color
var maxCubes = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

// checkSetsSanity returns if the set is possible or not
func checkSetsSanity(cubesSets string) (bool, error) {
	for _, cubesSet := range strings.Split(cubesSets, "; ") {
		for _, cubes := range strings.Split(cubesSet, ", ") {
			// Get number of cubes by color
			numberString := strings.Split(cubes, " ")[0]
			color := strings.Split(cubes, " ")[1]

			// Evaluate if the amount is possible
			if number, err := strconv.Atoi(numberString); err == nil {
				if number > maxCubes[color] {
					return false, nil
				}
			} else {
				return false, fmt.Errorf("number of cubes %s is not an integer", numberString)
			}
		}
	}

	return true, nil
}

// FirstPuzzle returns the result of the first puzzle
func FirstPuzzle() int {
	result := 0

	for _, inputLine := range strings.Split(Input, "\n") {
		cubesSets := strings.Split(inputLine, ": ")[1]
		if setResult, err := checkSetsSanity(cubesSets); err == nil {
			if setResult {
				gameIDString := strings.Split(strings.Split(inputLine, ": ")[0], " ")[1]
				if gameID, err := strconv.Atoi(gameIDString); err == nil {
					result += gameID
				} else {
					log.Fatal(err)
				}
			}
		} else {
			log.Fatal(err)
		}
	}

	return result
}

// getSetsPower calculates the powoer of sets
func getSetsPower(cubesSets string) (int, error) {
	setsPower := 1

	var minimumColors = map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	for _, cubesSet := range strings.Split(cubesSets, "; ") {
		for _, cubes := range strings.Split(cubesSet, ", ") {
			// Get number of cubes by color
			numberString := strings.Split(cubes, " ")[0]
			color := strings.Split(cubes, " ")[1]

			// Evaluate if the amount is possible and adapt if not
			if number, err := strconv.Atoi(numberString); err == nil {
				if number > minimumColors[color] {
					minimumColors[color] = number
				}
			} else {
				return -1, fmt.Errorf("number of cubes %s is not an integer", numberString)
			}
		}
	}

	// Calculate power
	for _, value := range minimumColors {
		setsPower *= value
	}

	return setsPower, nil
}

// Second puzzle returns the result of the second puzzle
func SecondPuzzle() int {
	result := 0

	for _, inputLine := range strings.Split(Input, "\n") {
		cubesSets := strings.Split(inputLine, ": ")[1]
		if setsPower, err := getSetsPower(cubesSets); err == nil {
			result += setsPower
		} else {
			log.Fatal(err)
		}
	}

	return result
}
