/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
    if head==nil || head.Next==nil{
        return head
    }
    var nh,current,t1,ret *ListNode
    ret= head.Next
    nh=head
    for nh.Next!=nil{
        current=nh.Next
        if current.Next!=nil{
            t1=current.Next
        }else{
            t1=nil
        }
        current.Next=nh
        if t1!=nil{
            if t1.Next==nil{
                nh.Next=t1
                break
            }
            nh.Next=t1.Next
        }else{
            nh.Next=t1
            break
        }
        nh=t1
    }
    if ret!=nil{
        return ret
    }else{
        return head
    }
}