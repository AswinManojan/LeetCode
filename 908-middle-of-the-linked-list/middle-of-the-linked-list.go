/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
    var p1,p2 *ListNode
    p1=head
    p2=head
    for p2!=nil{
        if p2.Next==nil{
            break
        }
        p2=p2.Next.Next
        p1=p1.Next
    }
    return p1
}