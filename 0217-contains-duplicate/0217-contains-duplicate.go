func containsDuplicate(nums []int) bool {
	freqMap := map[int]int{}
	for _, n := range nums {
		freqMap[n] += 1
		if freqMap[n] > 1 {
			return true
		}
	}
	return false
}