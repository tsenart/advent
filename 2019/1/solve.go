package main

import "fmt"

func main() {
	var sum int64
	for {
		var mass int64
		fmt.Scanf("%d", &mass)
		if mass == 0 {
			break
		}
		sum += fuel(mass)
	}
	fmt.Println(sum)
}

func fuel(mass int64) (sum int64) {
	for {
		fuel := mass/3 - 2

		if fuel <= 0 {
			return sum
		}

		mass = fuel
		sum += fuel
	}
}
