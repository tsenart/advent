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

	var p prog
	for sc.Scan() {
		fields := strings.Fields(sc.Text())
		ins := fields[0]
		arg, _ := strconv.Atoi(fields[1])
		p = append(p, op{ins, arg})
	}

	acc, term := p.run()

	fmt.Println("Part 1: ", acc)

	var mp int // mutation pointer
	for mp < len(p) {
		op := p[mp]
		mop := op // mutated op

		switch mop.ins {
		case "nop":
			mop.ins = "jmp"
		case "jmp":
			mop.ins = "nop"
		default:
			mp++
			continue
		}

		p[mp] = mop

		acc, term = p.run()
		if term {
			break
		}

		p[mp] = op // restore original op
		mp++
	}

	fmt.Println("Part 2: ", acc)
}

type op struct {
	ins string
	arg int
}

type prog []op

func (p prog) run() (acc int, term bool) {
	ip := 0
	exe := make([]int, len(p))
	for ip < len(p) {
		op := p[ip]

		if exe[ip]++; exe[ip] > 1 {
			return acc, false
		}

		switch op.ins {
		case "acc":
			acc += op.arg
			ip++
		case "jmp":
			ip += op.arg
		case "nop":
			ip++
		}
	}

	return acc, true
}
