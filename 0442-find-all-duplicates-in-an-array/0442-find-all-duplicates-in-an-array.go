func findDuplicates(nums []int) []int {
    array := []int{}
	myMap := map[int]int{}

	for _, num := range nums {
		myMap[num] += 1
		if myMap[num] == 2 {
			array = append(array, num)
		}
	}

	return array
}