func rearrangeArray(nums []int) []int {
    arr:= make([]int,len(nums))
    neg,pos:= 1,0
    for i:=0;i<len(nums);i++{
       if nums[i]>0{
           arr[pos]=nums[i]
           pos+=2
       }else{
           arr[neg]=nums[i]
           neg+=2
       }
    }
    return arr
}