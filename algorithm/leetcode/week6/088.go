/*
	88. Merge Sorted Array
	因为要把答案存到第一个数组中，所以会有覆盖问题，所以要从大到小排

	双指针（或者叫三指针）
*/
func merge(nums1 []int, m int, nums2 []int, n int)  {
	func merge(nums1 []int, m int, nums2 []int, n int)  {
		i,j,k := m-1,n-1,m+n-1
		for i>=0 && j>=0{
			if nums1[i]>nums2[j]{
				k-=1
				i-=1
				nums1[k] = nums1[i]
			}else{
				k-=1
				j-=1
				// todos::这里会数组越界
				nums1[k] = nums2[j]
			}
		}
		for j>=0{
			k-=1
			j-=1
			// todos::这里会数组越界
			nums1[k] = nums2[j]
		}
	}
}