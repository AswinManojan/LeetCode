func isPossibleToSplit(nums []int) bool {
    mp:=make(map[int]int)
    for i:=0;i<len(nums);i++{
        mp[nums[i]]++
    }
    for _,v:=range(mp){
        if v>2{
            return false
        }
    }
    return true
}