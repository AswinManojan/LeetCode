func pivotInteger(n int) int {
    s1,s2:=0,n
    for i,j:=1,n;i<=j;i++{
        s1+=i
        if s1>s2{
            j--
            s2+=j
        }
        if s1==s2 && i==j{
            return i
        }
    }
    return -1
}