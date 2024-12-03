package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 1 {
		log.Fatal("missing argument")
	}

	switch os.Args[1] {
	case "1":
		one(os.Stdin)
	case "2":
		two(os.Stdin)
	default:
		log.Fatal("invalid argument")
	}
}

func one(in io.Reader) {
	var (
		sc  = bufio.NewScanner(in)
		sum int
		mul int
	)

	for sc.Scan() {
		line := sc.Text()
		for line != "" {
			mul, line = evalMul(line)
			sum += mul
		}
	}

	log.Println(sum)
}

func evalMul(line string) (int, string) {
	_, after, ok := strings.Cut(line, "mul(")
	if !ok {
		return 0, ""
	}
	line = after

	between, after, ok := strings.Cut(after, ")")
	if !ok {
		return 0, line
	}

	first, second, ok := strings.Cut(between, ",")
	if !ok {
		return 0, line
	}

	fst, err := strconv.Atoi(first)
	if err != nil {
		return 0, line
	}

	snd, err := strconv.Atoi(second)
	if err != nil {
		return 0, line
	}

	return fst * snd, line
}

func two(in io.Reader) {
	var (
		sc  = bufio.NewScanner(in)
		sum int
		mul int
		on  = true
	)

	for sc.Scan() {
		line := sc.Text()
		for line != "" {
			if !on {
				_, after, ok := strings.Cut(line, "do()")
				line, on = after, ok
				continue
			}

			before, after, ok := strings.Cut(line, "don't()")
			for before != "" {
				mul, before = evalMul(before)
				sum += mul
			}
			line, on = after, !ok
		}
	}

	log.Println(sum)
}
