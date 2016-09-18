func removeDuplicates(nums []int) int {
    if  nums == nil || len(nums) == 0 {
        return 0
    }
    
    nowNum := nums[0]
    pos := 1
    
    for i:=1; i<len(nums); i++{
        if nums[i] == nowNum {
            continue
        }else{
            nowNum = nums[i]
            nums[pos] = nowNum
            pos++
        }
    }
    return pos

}
