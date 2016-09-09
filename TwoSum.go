func twoSum(nums []int, target int) []int {
    mapper := map[int]int{}
    var result []int
    
    for i := 0; i < len(nums); i ++{
        complement := target - nums[i]
        index, ok := mapper[complement]
        if (ok) {
            result = append (result, i)
            result = append (result, index)
            return result
        }
        mapper[nums[i]] = i
    }
    return nil
}
