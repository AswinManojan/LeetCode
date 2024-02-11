func maxSum(nums []int) int {
    max:=-1
    sum:=-1
    for i:=0;i<len(nums);i++{
        mi:=maxDigit(nums[i])
        for j:=i+1;j<len(nums);j++{
            if mi==maxDigit(nums[j]){
                sum=nums[i]+nums[j]
            }
            if sum>max{
                max=sum
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