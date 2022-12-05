package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	var stacks [][]byte
	for sc.Scan() {
		switch line := sc.Text(); {
		case line == "": // Parsing stacks done, reverse them.
			for stack := range stacks {
				for i := 0; i < len(stacks[stack])/2; i++ {
					j := len(stacks[stack]) - 1 - i
					stacks[stack][i], stacks[stack][j] = stacks[stack][j], stacks[stack][i]
				}
			}

		case strings.HasPrefix(line, "move"): // Move stack items.
			fields := strings.Fields(line[5:])
			moves, _ := strconv.Atoi(fields[0])
			from, _ := strconv.Atoi(fields[2])
			to, _ := strconv.Atoi(fields[4])

			from-- // zero indexing
			to--

			for moves > 0 {
				stacks[to] = append(stacks[to], stacks[from][len(stacks[from])-1])
				stacks[from] = stacks[from][:len(stacks[from])-1]
				moves--
			}

		default: // Parse stacks
			stack := 0
			for {
				if stack > len(stacks)-1 {
					stacks = append(stacks, []byte(nil))
				}

				if line[0] == '[' {
					stacks[stack] = append(stacks[stack], line[1])
				}

				if line = line[3:]; len(line) == 0 {
					break
				}

				line = line[1:] // Consume space
				stack++
			}
		}
	}

	for _, stack := range stacks {
		fmt.Printf("%c", stack[len(stack)-1])
	}
}
