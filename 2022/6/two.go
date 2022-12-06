package main

import (
	"fmt"
	"math/bits"
)

func main() {
	var datastream []byte
	fmt.Scanln(&datastream)

	for i := 0; i < len(datastream); i++ {
		var marker uint32
		for j := 0; j < 14; j++ {
			marker |= 1 << (datastream[i+j] - 'a')
		}

		if bits.OnesCount32(marker) == 14 {
			fmt.Println(i + 14)
			break
		}
	}
}
