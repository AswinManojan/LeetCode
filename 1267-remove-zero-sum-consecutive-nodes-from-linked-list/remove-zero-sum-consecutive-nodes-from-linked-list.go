/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeZeroSumSublists(head *ListNode) *ListNode {
    var prev *ListNode
    curr:=head
    if head.Val==0{
        head=head.Next
    }
    for curr!=nil{
        if prev!=nil && curr.Val==0{
            prev.Next=curr.Next
        }
        sum:=curr.Val
            c2:=curr.Next
        for c2!=nil{
            sum+=c2.Val
            if sum==0{
                if prev==nil{
                    head=c2.Next
                }else{
                    prev.Next=c2.Next
                }
            }
            c2=c2.Next
        }
        prev=curr
        curr=curr.Next
    }
    return head
}