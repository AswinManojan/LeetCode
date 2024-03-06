func maxSubArray(nums []int) int {
    sum,max:=0,-1*int(math.Pow(10,4))
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