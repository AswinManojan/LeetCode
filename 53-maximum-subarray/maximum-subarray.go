func maxSubArray(nums []int) int {
    sum,max:=0,nums[0]
    for i:=0;i<len(nums);i++{
        sum+=nums[i]
        if sum>max{
            max=sum
        }
        if sum<0{
            sum=0
        }
    }
    return max
}