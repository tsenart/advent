package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	p := make(passport)
	solutions := [2]int{}
	scan := true

	for scan {
		if scan = sc.Scan(); scan && sc.Text() != "" {
			p.Parse(sc.Text())
			continue
		}

		if p.Valid() {
			solutions[0]++
		}

		if p.ValidStrict() {
			solutions[1]++
		}

		p.Reset()
	}

	fmt.Println(solutions)
}

type passport map[string]string

func (p passport) Parse(line string) {
	for _, field := range strings.Fields(line) {
		ps := strings.SplitN(field, ":", 2)
		p[ps[0]] = ps[1]
	}
}

func (p passport) Valid() bool {
	for _, field := range []string{
		"byr",
		"iyr",
		"eyr",
		"hgt",
		"hcl",
		"ecl",
		"pid",
		// 	"cid"
	} {
		if v, ok := p[field]; !ok || v == "" {
			return false
		}
	}
	return true
}

func (p passport) ValidStrict() bool {
	byr, _ := strconv.Atoi(p["byr"])
	if byr < 1920 || byr > 2002 {
		return false
	}

	iyr, _ := strconv.Atoi(p["iyr"])
	if iyr < 2010 || iyr > 2020 {
		return false
	}

	eyr, _ := strconv.Atoi(p["eyr"])
	if eyr < 2020 || eyr > 2030 {
		return false
	}

	hgt := p["hgt"]
	if len(hgt) < 2 {
		return false
	}

	unit := hgt[len(hgt)-2:]
	height, _ := strconv.Atoi(hgt[:len(hgt)-2])

	switch unit {
	case "in":
		if height < 56 || height > 76 {
			return false
		}
	case "cm":
		if height < 150 || height > 193 {
			return false
		}
	default:
		return false
	}

	if ok, _ := regexp.MatchString(`^#[0-9a-f]{6}$`, p["hcl"]); !ok {
		return false
	}

	switch p["ecl"] {
	case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
	default:
		return false
	}

	if ok, _ := regexp.MatchString(`^[0-9]{9}$`, p["pid"]); !ok {
		return false
	}

	return true
}

func (p passport) Reset() {
	for k := range p {
		delete(p, k)
	}
}
