func sortedSquares(nums []int) []int {
	l, r := 0, len(nums) -1
	array := []int{}
	for i := 0; i < len(nums); i++{
		if math.Pow(float64(nums[r]),2) > math.Pow(float64(nums[l]),2){
			array = append(array, int(math.Pow(float64(nums[r]),2)))
			r -= 1
		}else{
			array = append(array, int(math.Pow(float64(nums[l]), 2)))
			l +=1
		}
	}
	return reverseSlice(array)
}

func reverseSlice(nums []int) []int {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	return nums
}