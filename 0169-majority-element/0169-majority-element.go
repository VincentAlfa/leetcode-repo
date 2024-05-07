func majorityElement(nums []int) int {
    hashMap := map[int]int{}
	max := 0
	for _, value := range nums {
		hashMap[value] += 1
		if hashMap[value] > max {
			max = hashMap[value]
		}
	}
	
	for key := range hashMap {
		if hashMap[key] == max {
			return key
		}
	}

	return 0
}