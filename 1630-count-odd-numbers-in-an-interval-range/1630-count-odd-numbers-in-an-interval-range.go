func countOdds(low int, high int) int {
	counter := 0

	for i := low; i <= high; i++ {
		if i%2 != 0 {
			counter++
		}
	}

	return counter
}