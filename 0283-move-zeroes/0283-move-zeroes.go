func moveZeroes(nums []int) {
	l := 0

	for r := range nums {
		if nums[r] != 0 {
			temp := nums[l]
			nums[l] = nums[r]
			nums[r] = temp
			l++
		}
		
	}
}