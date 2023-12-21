package main

import (
	"fmt"
	dayOne "puzzles/puzzles/1st"
	dayTwo "puzzles/puzzles/2nd"
	dayThree "puzzles/puzzles/3rd"
	dayFour "puzzles/puzzles/4th"
	dayFive "puzzles/puzzles/5th"
)

func main() {
	fmt.Println("1st December, first puzzle solution :", dayOne.FirstPuzzle())
	fmt.Println("1st December, second puzzle solution :", dayOne.SecondPuzzle())
	fmt.Println("2nd December, first puzzle solution :", dayTwo.FirstPuzzle())
	fmt.Println("2nd December, second puzzle solution :", dayTwo.SecondPuzzle())
	fmt.Println("3rd December, first puzzle solution :", dayThree.FirstPuzzle())
	// fmt.Println("3rd December, second puzzle solution :", dayThree.SecondPuzzle())
	fmt.Println("4th December, first puzzle solution :", dayFour.FirstPuzzle())
	fmt.Println("4th December, second puzzle solution :", dayFour.SecondPuzzle())
	fmt.Println("5th December, first puzzle solution :", dayFive.FirstPuzzle())
	fmt.Println("5th December, second puzzle solution :", dayFive.SecondPuzzle())
}
