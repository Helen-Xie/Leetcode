func merge(nums1 []int, m int, nums2 []int, n int)  {
    var f1, f2 int
    
    if n == 0 {
        return
    }
    
    if m == 0 {
        for i, _ := range nums2  {
            nums1[i] = nums2[i]
        }
        return
    }
    
    if nums1[0]<=nums1[m-1]{
        f1=0
    }else{
        f1=1
    }
    
    if nums2[0]<=nums2[n-1]{
        f2=0
    }else{
        f2=1
    }
    
    if f1!= f2 {
        for i:=0; i<=n/2; i++ {
            nums2[i] = nums2[n-1-i]
        }
    }
    
    pos1 := m-1
    pos2 := n-1
    pos := m+n-1
    
    if f1 != 0  {
        for pos1>=0 && pos2 >=0 {
            if nums1[pos1]<nums2[pos2] {
                nums1[pos] = nums1[pos1]
                pos1--
                pos--
            } else {
                nums1[pos] = nums2[pos2]
                pos2--
                pos--
            }
        }
        if(pos2>=0){
                for i:=0;i<=pos2;i++ {
                    nums1[i]=nums2[i];
            }
        }
    } else {
        for pos1>=0 && pos2 >=0 {
            if nums1[pos1]>nums2[pos2] {
                nums1[pos] = nums1[pos1]
                pos1--
                pos--
            } else {
                nums1[pos] = nums2[pos2]
                pos2--
                pos--
            }
        } 
        if(pos2>=0){
            for i:=0;i<=pos2;i++ {
                nums1[i]=nums2[i];
            }
        }
    }
    
    return
}
