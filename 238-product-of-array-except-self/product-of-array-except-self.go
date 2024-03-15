func productExceptSelf(nums []int) []int {
    r1:=make([]int,len(nums))
    curr:=1
    for i:=0;i<len(nums);i++{
        r1[i]=curr
        curr*=nums[i]
        // r1=append(r1,nums[i]*r1[i])
    }
    curr=1
    for i:=len(nums)-1;i>=0;i--{
        r1[i]*=curr
        curr*=nums[i]
    }
    return r1
}