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

		switch round {
		case "A X", "B Y", "C Z": // Rock Rock, Paper Paper, Scissors Scissors
			score += 3
		case "A Y", "B Z", "C X": // Rock Paper, Paper Scissors, Scissors Rock
			score += 6
		}

		switch round[2:] {
		case "X": // Loose
			score += 1
		case "Y": // Draw
			score += 2
		case "Z": // Win
			score += 3
		}
	}

	fmt.Println(score)
}
