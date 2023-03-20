package collatzconjecture

import "fmt"

func collate(n, counter int) int {
	if n == 1 {
		return counter
	}

	if n%2 == 0 {
		return collate(n/2, counter+1)
	}

	return collate(3*n+1, counter+1)
}

func CollatzConjecture(n int) (int, error) {
	if n < 1 {
		return 0, fmt.Errorf("invalid input")
	}

	return collate(n, 0), nil
}
