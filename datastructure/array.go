package datastructure

func merge(nums1 []int, m int, nums2 []int, n int)  {
    var j,k int
    var temp  = make([]int,m)
	for i := range nums1[:m]{
		temp[i] = nums1[i]
	}
    
    for i:=0;i<m+n && j < m && k < n;i++ {
        if temp[j] < nums2[k] {
            nums1[i] = temp[j]
            j++
        }else {
            nums1[i] = nums2[k]
            k++
        }
    }
    
    if j>=m {
		for i := j+1; i < m+n; i++ {
			nums1[i] = nums2[k]
			k++
		}
		return
    } 
        
    if k>=n {
		for i := k+1; i < m+n; i++ {
			nums1[i] = temp[j]
			j++
		}
    }
    
}