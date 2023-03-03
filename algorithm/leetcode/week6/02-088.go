/*
	方法1：双指针解法
*/
func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, m+n-1
	// question： 从后向前遍历（因为要把答案存到第一个数组中，所以会有覆盖问题，所以要从大到小排）
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			k--
			i--
		} else {
			nums1[k] = nums2[j]
			k--
			j--
		}
	}
	// 因为要把答案存到第一个数组中,如果i有剩余，那么已经在正确的位置，就不需要调整了
	for j >= 0 {
		nums1[k] = nums2[j]
		k--
		j--
	}
}

/*
	思路（还是上一个思路简单）：
		1. 遍历两个数组，把符合条件的先挪到一起
		2. 剩下的部分直接合并
*/
func merge(nums1 []int, m int, nums2 []int, n int) {
	res := make([]int, 0)
	i, j := 0, 0
	for i < m && j < n {
		if nums1[i] < nums2[j] {
			res = append(res, nums1[i])
			i++
		} else {
			res = append(res, nums2[j])
			j++
		}
	}
	if i < m {
		res = append(res, nums1[i:]...)
	}
	if j < n {
		res = append(res, nums2[j:]...)
	}
	nums1 = res
}

