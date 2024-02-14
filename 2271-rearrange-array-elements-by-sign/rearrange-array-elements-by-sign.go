func rearrangeArray(nums []int) []int {
    var arr []int
    var i,j=0,0
    for k:=0;k<len(nums)/2;k++{
        for ;i<len(nums);i++{
            if nums[i]>0{
                arr=append(arr,nums[i])
                i++
                break
            }
        }
        for ;j<len(nums);j++{
            if nums[j]<0{
                arr=append(arr,nums[j])
                j++
                break
            }
        }
    }
    return arr
}