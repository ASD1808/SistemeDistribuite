package main

import (
	"fmt"
)

func main() {
	sirNumere := []int32{10, 15, 22, 33, 40, 51, 64}

	var numerePare []int32

	for _, num := range sirNumere {
		if num%2 == 0 {
			numerePare = append(numerePare, num)
		}
	}

	fmt.Println("Numerele pare sunt:", numerePare)
}
