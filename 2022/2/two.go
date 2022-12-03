package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var score int

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		round := sc.Text()

		switch round[2:] {
		case "Y": // Draw
			score += 3
		case "Z": // Win
			score += 6
		}

		switch round {
		case "C Z", "A Y", "B X": // Scissors Rock, Rock Rock, Paper Rock
			score += 1
		case "A Z", "B Y", "C X": // Rock Paper, Paper Paper, Scissors Paper
			score += 2
		case "B Z", "C Y", "A X": // Paper Scissors, Scissors Scissors, Rock Scissors
			score += 3
		}
	}

	fmt.Println(score)
}
