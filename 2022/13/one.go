package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	pair := make([]packet, 0, 2)
	index := 1
	sum := 0

	for sc.Scan() || len(pair) != 0 {
		if len(pair) == 2 {
			if pair[0].Compare(pair[1]) < 0 {
				sum += index
			}
			pair = pair[:0]
			index++
		}

		if line := sc.Bytes(); len(line) != 0 {
			_, pkt := parse(line)
			pair = append(pair, pkt)
		}
	}

	fmt.Println(sum)
}

type packet interface {
	Compare(packet) int
}

func parse(p []byte) ([]byte, packet) {
	if len(p) == 0 {
		return p, nil
	}

	if p[0] != '[' {
		i := bytes.IndexAny(p, ",]")
		if i <= 0 {
			return p, nil
		}

		n, _ := strconv.Atoi(string(p[:i]))
		return p[i:], integer(n)
	}

	var l list
	for p[0] != ']' {
		var pkt packet
		if p, pkt = parse(p[1:]); pkt != nil {
			l = append(l, pkt)
		}
	}

	return p[1:], l
}

type list []packet

func (l list) Compare(other packet) int {
	o, ok := other.(list)
	if !ok {
		o = list{other}
	}

	for i, p := range l {
		if i == len(o) {
			break
		} else if c := p.Compare(o[i]); c != 0 {
			return c
		}
	}

	return len(l) - len(o)
}

type integer int

func (i integer) Compare(other packet) int {
	if o, ok := other.(integer); ok {
		return int(i - o)
	}
	return list{i}.Compare(other)
}
