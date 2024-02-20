func wateringPlants(plants []int, capacity int) int {
    can:=capacity
    count:=0
    for i:=0;i<len(plants);i++{
        if plants[i]<=can{
            can-=plants[i]
            count++
        }else{
            count+=(2*i)
            can=capacity
            i--
        }
    }
    return count
}