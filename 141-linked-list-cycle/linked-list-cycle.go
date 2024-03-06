/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    mp:=make(map[*ListNode]bool)
    current:=head
    for current!=nil{
        if !mp[current]{
            mp[current]=true
        }else{
            return true
        }
        current=current.Next
    }
    return false
}