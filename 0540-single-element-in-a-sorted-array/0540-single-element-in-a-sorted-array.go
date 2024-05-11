func singleNonDuplicate(nums []int) int {
	hashmap := map[int]int{}

	for _, value := range nums {
		hashmap[value] += 1
	}

	for key, value := range hashmap {
		if value == 1 {
			return key
		}
	}
	return 0
}