package dayFive

import (
	"log"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func getSeeds(lineValues string) []int {
	seeds := []int{}
	re := regexp.MustCompile("[0-9]+")

	for _, value := range re.FindAllString(lineValues, -1) {
		if seed, err := strconv.Atoi(value); err == nil {
			seeds = append(seeds, seed)
		} else {
			log.Fatal(err)
		}
	}

	return seeds
}

func getSeedResut(result int, data []string) int {
	for _, val := range data {
		re := regexp.MustCompile("[0-9]+")
		sourceToDestination := re.FindAllString(val, -1)

		// Get source/destination information in line
		destinationStart, err := strconv.Atoi(sourceToDestination[0])
		if err != nil {
			log.Fatal(err)
		}
		sourceStart, err := strconv.Atoi(sourceToDestination[1])
		if err != nil {
			log.Fatal(err)
		}
		rangeLength, err := strconv.Atoi(sourceToDestination[2])
		if err != nil {
			log.Fatal(err)
		}

		if result >= sourceStart && result < sourceStart+rangeLength {
			seedRange := result - sourceStart
			return destinationStart + seedRange
		}
	}

	return result
}

// FirstPuzzle returns the result of the first puzzle
func FirstPuzzle() int {
	resultPerSeed := []int{}
	sourceToDestination := []string{}

	for lineIndex, line := range strings.Split(Input, "\n") {
		if lineIndex == 0 {
			// Initialize result per seed
			resultPerSeed = getSeeds(strings.Split(line, ": ")[1])

		} else if lineIndex == 1 {
			continue

		} else if line == "" {
			// Update result for each seed depending on source/destination information
			for index, value := range resultPerSeed {
				resultPerSeed[index] = getSeedResut(value, sourceToDestination)
			}
			sourceToDestination = []string{}

		} else if !strings.Contains(line, "map") { // Update result per sourc
			sourceToDestination = append(sourceToDestination, line)
		}
	}

	// Update result for each seed depending on source/destination information
	for index, value := range resultPerSeed {
		resultPerSeed[index] = getSeedResut(value, sourceToDestination)
	}

	return slices.Min(resultPerSeed)
}

func getMoreSeeds(lineValues string) []int {
	seeds := []int{}
	re := regexp.MustCompile("[0-9]+")

	var startSeed int
	for index, value := range re.FindAllString(lineValues, -1) {
		if val, err := strconv.Atoi(value); err == nil {
			if index%2 == 0 {
				startSeed = val
			} else {
				for i := 0; i < val; i++ {
					seeds = append(seeds, startSeed+i)
				}
			}
		} else {
			log.Fatal(err)
		}
	}

	return seeds
}

func SecondPuzzle() int {
	resultPerSeed := []int{}
	sourceToDestination := []string{}

	for lineIndex, line := range strings.Split(Input, "\n") {
		if lineIndex == 0 {
			// Initialize result per seed
			resultPerSeed = getMoreSeeds(strings.Split(line, ": ")[1])

		} else if lineIndex == 1 {
			continue

		} else if line == "" {
			// Update result for each seed depending on source/destination information
			for index, value := range resultPerSeed {
				resultPerSeed[index] = getSeedResut(value, sourceToDestination)
			}
			sourceToDestination = []string{}

		} else if !strings.Contains(line, "map") { // Update result per sourc
			sourceToDestination = append(sourceToDestination, line)
		}
	}

	// Update result for each seed depending on source/destination information
	for index, value := range resultPerSeed {
		resultPerSeed[index] = getSeedResut(value, sourceToDestination)
	}

	return slices.Min(resultPerSeed)
}
