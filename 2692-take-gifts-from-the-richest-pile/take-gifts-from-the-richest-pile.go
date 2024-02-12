import "math"
func pickGifts(gifts []int, k int) int64 {
    for i:=0;i<k;i++{
        sort.Ints(gifts)
        gifts[len(gifts)-1]=sqRt(gifts[len(gifts)-1])
    }
    sum:=0
    for i:=0;i<len(gifts);i++{
        sum+=gifts[i]
    }
    return int64(sum)
}

func sqRt(num int) int{
    return int(math.Sqrt(float64(num)))
}