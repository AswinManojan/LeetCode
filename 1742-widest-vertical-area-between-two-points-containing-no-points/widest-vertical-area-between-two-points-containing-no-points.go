func maxWidthOfVerticalArea(points [][]int) int {
    var arr []int
    for i:=0;i<len(points);i++{
        arr=append(arr,points[i][0])
    }
    sort.Ints(arr)
    max:=0
    for i:=1;i<len(arr);i++{
        if max<(arr[i]-arr[i-1]){
            max=arr[i]-arr[i-1]
        }
    }
    return max
}