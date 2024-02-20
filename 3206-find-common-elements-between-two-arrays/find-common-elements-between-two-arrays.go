func findIntersectionValues(nums1 []int, nums2 []int) []int {
    m1,m2:=make(map[int]bool),make(map[int]bool)
    for i:=0;i<len(nums1);i++{
        m1[nums1[i]]=true
    }
    for i:=0;i<len(nums2);i++{
        m2[nums2[i]]=true
    }
    var c1,c2 int
    arr:= make([]int,2)
    for i:=0;i<len(nums1);i++{
        if m2[nums1[i]]{
            c1++
        }
    }
    for i:=0;i<len(nums2);i++{
        if m1[nums2[i]]{
            c2++
        }
    }
    arr[0]=c1
    arr[1]=c2
    return arr
}