package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Initialize the reader pointer
	reader := bufio.NewReader(os.Stdin)

	// Read the number of test cases
	var T int
	fmt.Scanf("%d", &T)

	// Iterate for each test case
	for t := 0; t < T; t++ {
		// Read the number of pancakes
		pancakes, _ := reader.ReadString('\n')

		// Hold the number of shifts and the current sign
		answer, sign := 0, pancakes[0]

		// Iterate over each pancake index
		for i := 0; i < len(pancakes) - 1; i++ {
			if sign != pancakes[i] {
				sign = pancakes[i]
				answer++
			}
		}

		// If the current sign is '-', shift them one more time
		if sign == '-' {
			answer++
		}

		// Print the minimum number of shifts
		fmt.Printf("Case #%d: %d\n", t + 1, answer)
	}
}