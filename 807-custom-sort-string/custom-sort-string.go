func customSortString(order string, s string) string {
    res:=""
    mp:=make(map[rune]int)
    for _,v:= range(s){
        mp[v]++
    }
    for _,v:=range(order){
        for mp[v]>0{
            res+=string(v)
            mp[v]--
        }
    }
    for _,v:=range(s){
        if mp[v]>0{
            mp[v]--
            res+=string(v)
        }
    }
    return res
}