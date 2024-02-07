func findIndices(nums []int, indexDifference int, valueDifference int) []int {
    result:=[]int{-1,-1}
    for i:=0;i<len(nums);i++{
        if i+indexDifference<len(nums){
            for j:=i+indexDifference;j<len(nums);j++{
                if abs(nums[i] - nums[j]) >= valueDifference{
                    result[0]=i
                    result[1]=j
                    return result
                }
            }
        }
    }
    return result
}

func abs(i int) int{
    if 0>i{
        return -i
    }
    return i
}