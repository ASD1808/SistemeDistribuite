package main

import (
	"fmt"
)

func main() {
	var intArr = [5]int32{42, 15, 8, 23, 4}
	n := len(intArr)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if intArr[j] > intArr[j+1] {
				temp := intArr[j]
				intArr[j] = intArr[j+1]
				intArr[j+1] = temp
			}
		}
	}

	fmt.Println("Tabloul sortat crescător:", intArr)
}
