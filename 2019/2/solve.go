package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	prog, err := parse(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	solution1, err := solve1(prog)
	if err != nil {
		log.Fatal(err)
	}

	solution2, err := solve2(prog)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("solution 1 = %d\n", solution1)
	fmt.Printf("solution 2 = %d\n", solution2)
}

func parse(r io.Reader) ([]int, error) {
	input, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	ops := strings.Split(string(input), ",")
	prog := make([]int, len(ops))
	for i, op := range ops {
		prog[i], _ = strconv.Atoi(op)
	}

	return prog, nil
}

func run(prog []int) error {
	for ip := 0; prog[ip] != 99; ip += 4 {
		a, b, c := prog[prog[ip+1]], prog[prog[ip+2]], &prog[prog[ip+3]]
		switch op := prog[ip]; op {
		case 1:
			*c = a + b
		case 2:
			*c = a * b
		default:
			return fmt.Errorf("unknown opcode prog[%d]=%d", ip, op)
		}
	}
	return nil
}

func solve1(prog []int) (int, error) {
	dup := append(prog[:0:0], prog...)
	dup[1], dup[2] = 12, 2 // Inputs
	err := run(dup)
	return dup[0], err
}

func solve2(prog []int) (int, error) {
	dup := make([]int, len(prog))
	noun, verb := 0, 0

	for {
		copy(dup, prog)
		dup[1], dup[2] = noun, verb

		switch err := run(dup); {
		case err != nil:
			return 0, err
		case dup[0] == 19690720:
			return 100*noun + verb, nil
		case noun < 99:
			noun++
		case verb < 99:
			noun = 0
			verb++
		default:
			return 0, fmt.Errorf("no solution found")
		}
	}
}
