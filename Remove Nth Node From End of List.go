func removeNthFromEnd(head *ListNode, n int) *ListNode {
    if n == 0 {
        return head
    }
    
    p := head
    pre := head
    
    for i := 0; i < n; i++{
        p=p.Next
    }
    if p == nil {
        head = head.Next
        return head
    }

    for (p.Next != nil ) {
        pre=pre.Next
        p=p.Next
    }
    
    pre.Next = pre.Next.Next
    return head
}

