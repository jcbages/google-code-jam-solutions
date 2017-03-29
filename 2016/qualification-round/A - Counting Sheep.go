package main

import "fmt"

func main() {
	// Read the number of test cases
	var T int
	fmt.Scanf("%d", &T)

	// Iterate for each test case
	for t := 0; t < T; t++ {
		// Read the initial number N
		var N int
		fmt.Scanf("%d", &N)

		// Base case, N is zero so counts forever
		if N == 0 {
			fmt.Printf("Case #%d: INSOMNIA\n", t + 1)
			continue
		}

		// Prepare the marked boolean bits and increment
		marked, increment := 0, N

		// Count until every digit has been found
		for ; marked < (1 << 10) - 1; N += increment {
			// Mark each digit as counter
			for M := N; M > 0; M /= 10 {
				marked |= (1 << uint(M % 10))
			}
		}

		// Print the last number counter
		fmt.Printf("Case #%d: %d\n", t + 1, N - increment)
	}
}