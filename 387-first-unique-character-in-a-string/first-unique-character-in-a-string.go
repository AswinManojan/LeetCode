func firstUniqChar(s string) int {
    mp:=make(map[rune]int)
    for _,x:=range(s){
        mp[x]++
    }
    for i,x:=range(s){
        if mp[x]==1{
            return i
        }
    }
    return -1
}