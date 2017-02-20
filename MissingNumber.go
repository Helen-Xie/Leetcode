func missingNumber(nums []int) int {
	var i,t int
	for i=0; i<len(nums); i++ {
		for nums[i]!=i && nums[i]!=len(nums) {
			t=nums[nums[i]]
			nums[nums[i]]=nums[i]
			nums[i]=t
		}
	}
	for i=0; i<len(nums); i++ {
		if(nums[i]!=i){
			return i
		}
	}

	return len(nums)
}