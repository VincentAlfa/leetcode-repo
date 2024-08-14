function countOdds(low: number, high: number): number {
    let counter: number = 0

    for (let i = low; i <= high; i++) {
		if (i%2 != 0) {
			counter++
		}
	}
    return counter
};