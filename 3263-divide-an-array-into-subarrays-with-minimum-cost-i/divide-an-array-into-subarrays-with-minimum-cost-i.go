func minimumCost(nums []int) int {
    first:=nums[0]
    arr:=nums[1:]
    sort.Ints(arr)
    return first+arr[0]+arr[1]
}