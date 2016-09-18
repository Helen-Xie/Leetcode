/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil {
        return l2
    }else if l2 == nil {
        return l1
    }
    
    var val1[1000] int
    var val2[1000] int 
    var val[2000] int
    num1 := 0; num2 := 0
    p1 := l1 
    p2 := l2
    
    val1[num1] = p1.Val
    num1++
    for p1.Next != nil {
        val1[num1] = p1.Next.Val
        num1++
        p1 = p1.Next
    }
    
    val2[num2] = p2.Val
    num2++
    for p2.Next != nil {
        val2[num2] = p2.Next.Val
        num2++
        p2 = p2.Next
    }
    
    p1.Next = l2
    
    if val1[0] > val1[num1-1]{
        for i:=0; i<num1/2; i++{
            t:=val1[i]
            val1[i]=val1[num1-1-i]
            val1[num1-1-i]=t
        }
    }
    if val2[0]>val2[num2-1] {
        for i:=0;i<num2/2;i++ {
            t:=val2[i]
            val2[i]=val2[num2-1-i]
            val2[num2-1-i]=t
        }
    }
    
    num:=0
    j:=num
    i:=j
    
    for i<num1 && j<num2 {
        if(val1[i]<val2[j]){
            val[num]=val1[i]
            i++
            num++
        }else{
            val[num]=val2[j]
            j++
            num++
        }
    }
    
    if(i<num1){
        for i<num1 {
            val[num]=val1[i]
            i++
            num++
        }
    }else{
        for j<num2 {
            val[num]=val2[j]
            j++
            num++
        }
    }
    
    p1=l1
    i=0
    for p1!=nil {
        p1.Val=val[i]
        i++
        p1=p1.Next
    }
    
    return l1
}
