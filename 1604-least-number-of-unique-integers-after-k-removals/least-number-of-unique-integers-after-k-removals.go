func findLeastNumOfUniqueInts(arr []int, k int) int {
    mp:=make(map[int]int)
    for i:=0;i<len(arr);i++{
        mp[arr[i]]++
    }
    res:=sortValue(mp)
    count:=0
    for i:=0;k>0;i++{
        if k>=mp[res[i]]{
            k-=mp[res[i]]
            count++
        }else{
            break
        }
    }
    return len(mp)-count
}
func sortValue (mp map[int]int) []int{
    var sortedKeys []int
    for keys:= range mp{
        sortedKeys=append(sortedKeys,keys)
    }
    sort.Slice(sortedKeys,func(i,j int) bool{
        return mp[sortedKeys[i]]<mp[sortedKeys[j]]
    })
    return sortedKeys
}
