func removeElement(nums []int, val int) int {
    l := 0 
	for  r := range nums {
		if nums[r] != val {
			nums[l] = nums[r]
			l += 1
		}
	}
	return l 
}