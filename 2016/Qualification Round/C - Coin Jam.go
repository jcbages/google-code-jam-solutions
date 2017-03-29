package main

import (
	"fmt"
)

func main() {
	// Read the values to identify small/large case
	var T, N, J int
	fmt.Scanf("%d\n%d %d", &T, &N, &J)

	// Build the basic string to print [1, 0, ..., 0, 1]
	number := make([]byte, N - 1)
	number[0], number[N-2] = 1, 1

	// Generate and print J coin jams
	fmt.Println("Case #1:")
	for i := 2; i < N - 3; i++ {
		for j := i + 2; j < N - 3; j++ {
			for k := j + 2; k < N - 3; k++ {
				// Build the number by adding 1s on i, j, k
				number[i], number[j], number[k] = 1, 1, 1

				// Print the coin jam -> number + (number << 1)
				fmt.Print(1)
				for l := 1; l < N - 1; l++ {
					fmt.Print(number[l] + number[l-1])
				}
				fmt.Print(1)
				fmt.Println(" 3 4 5 6 7 8 9 10 11")

				// Restore the number array by adding 0s on i, j, k
				number[i], number[j], number[k] = 0, 0 , 0

				// Check if J numbers have been printed
				J -= 1
				if J == 0 {
					return
				}
			}
		}
	}
}
