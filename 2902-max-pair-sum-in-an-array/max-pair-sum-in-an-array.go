func maxSum(nums []int) int {
    max:=-1
    for i:=0;i<len(nums);i++{
        mi:=maxDigit(nums[i])
        for j:=i+1;j<len(nums);j++{
            if mi==maxDigit(nums[j]){
                if max<(nums[i]+nums[j]){
                    max=nums[i]+nums[j]
                }
            }
        }
    }
    return max
}
func maxDigit(num int) int{
    max:=0
    for num>0{
        if num%10>max{
            max=num%10
        }
        num/=10
    }
    return max
}