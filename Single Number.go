func singleNumber(nums []int) int {
	var t int
	t = nums[0]
	for i := 1; i < len(nums); i++ {
	t ^= nums[i]
	}

	return t
}
