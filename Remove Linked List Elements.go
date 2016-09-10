func removeElements(head *ListNode, val int) *ListNode {
    for head != nil && head.Val == val{
        head = head.Next
    }
    
    if head == nil {
        return nil
    }
    
    p := head
    for p.Next != nil {
        if p.Next.Val == val {
            p.Next = p.Next.Next
            continue
        }
        p = p.Next
    }
    
    return head
}
