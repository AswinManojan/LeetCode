func findJudge(n int, trust [][]int) int {
    mp:=make(map[int]int)
    mp2:=make(map[int]int)
    for i:=0;i<len(trust);i++{
        mp[trust[i][0]]++
        mp2[trust[i][1]]++
    }
    for i:=1;i<=n;i++{
        if mp[i]==0 && mp2[i]==n-1{
            return i
        }
    }
    return -1
}