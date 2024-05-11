function sortedSquares(nums: number[]): number[] {
    return nums.map((nums) => Math.pow(nums, 2)).sort((a,b) => a-b)
};