func maxFrequencyElements(nums []int) int {
    mp:=make(map[int]int)
    for _,v:=range(nums){
        mp[v]++
    }
    arr:=[]int{}
    for k,_:=range(mp){
        arr=append(arr,k)
    }
    sort.Slice(arr,func(i,j int)bool{
        return mp[arr[j]]<mp[arr[i]]
    })
    fmt.Println(arr)
    val:=mp[arr[0]]
    sum:=mp[arr[0]]
    for i:=1;i<len(arr);i++{
        if val==mp[arr[i]]{
            sum+=mp[arr[i]]
        }else{
            break
        }
    }
    return sum
}