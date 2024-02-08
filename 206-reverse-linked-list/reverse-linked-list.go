/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    current:=head
    var prev,nxt *ListNode
    for current!=nil{
        nxt=current.Next
        current.Next=prev
        prev=current
        current=nxt
    }
    return prev
}