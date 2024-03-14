func numSubarraysWithSum(nums []int, goal int) int {
    count,sum:=0,0
    for i:=0;i<len(nums);i++{
        sum=nums[i]
        if sum==goal{
            count++
        }
        for j:=i+1;j<len(nums)&&(sum<=goal);j++{
            sum+=nums[j]
            if sum==goal{
                count++
            }
        }
    }
    return count
}