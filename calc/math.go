package calc

// Add returns sum of integers
func Add(numbers ...int) int {
	sum := 0

	for _, num := range numbers {
		sum += num
	}

	return sum
}
