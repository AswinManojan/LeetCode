/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNodes(head *ListNode) *ListNode {
    head=reverseLL(head)
    temp:= &ListNode{0,head}
    curr:= head
    prev:= temp
    currMax:= head.Val
    for curr!=nil{
        if curr.Val<currMax{
            prev.Next=curr.Next
        }else{
            currMax=curr.Val
            prev=curr
        }
        curr =curr.Next
    }
    return reverseLL(temp.Next);
}

func reverseLL(head *ListNode) *ListNode{
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