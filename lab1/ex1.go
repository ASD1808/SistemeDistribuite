package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	
	scanner := bufio.NewScanner(os.Stdin)
	
	scanner.Scan()

	v := scanner.Text()

	x, err := strconv.Atoi(v)

	if err != nil {
		fmt.Println("Eroare la conversie:")
	}

	fmt.Printf("%T", x)

	if x <= 1 {
		fmt.Println("Numarul nu este prim")
	}
	for i := 2; i < x/2; i++ {
		if x%i == 0 {
			fmt.Println("Numarul nu este prim")
		}
	}
	fmt.Println("Numarul este prim")
}
