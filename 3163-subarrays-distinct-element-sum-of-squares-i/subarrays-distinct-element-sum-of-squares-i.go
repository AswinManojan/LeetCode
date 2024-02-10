func sumCounts(nums []int) int {
    var arr []int
    for i:=0;i<len(nums);i++{
        for j:=i;j<len(nums);j++{
            mp:=make(map[int]bool)
            for _,x:=range(nums[i:j+1]){
                mp[x]=true
            }
            arr=append(arr,len(mp))
        }
    }
    var res int
    for _,x:=range arr{
        res+=(x*x)
    }
    return res
}