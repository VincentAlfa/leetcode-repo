func twoSum(nums []int, target int) []int {
    hashMap := map[int]int{}

	for index, value := range nums {
		complement := target - value
		_, isExist := hashMap[complement]
		if isExist {
			return []int{hashMap[complement], index}
		}
		hashMap[value] = index
	}
	return []int{}
}