func differenceOfSums(n int, m int) int {
    var s1,s2 int
    for i:=1;i<=n;i++{
        if i%m!=0{
            s1+=i
        }else{
            s2+=i
        }
    }
    return s1-s2
}