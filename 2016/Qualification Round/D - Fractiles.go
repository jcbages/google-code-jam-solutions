package main

import (
	"fmt"
	"math"
)

func main() {
	// Read the number of test cases
	var T int
	fmt.Scanf("%d", &T)

	// Iterate for each test case
	for t := 0; t < T; t++ {
		// Read the necessary data
		var K, C, S int
		fmt.Scanf("%d %d %d", &K, &C, &S)

		// Calculate the chunk size = K^C / S
		chunk := int(math.Pow(float64(K), float64(C))) / S

		// Print the answer for the current case
		fmt.Printf("Case #%d:", t + 1)
		for i := 0; i < S; i++ {
			fmt.Printf(" %d", chunk * i + 1)
		}
		fmt.Println()
	}
}
