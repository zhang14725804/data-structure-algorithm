/*
	Merge Sorted Array
	给你两个有序整数数组 nums1 和 nums2，请你将 nums2 合并到 nums1 中，使 nums1 成为一个有序数组。
*/

/*
	因为要把答案存到第一个数组中，所以会有覆盖问题，所以要从大到小排（很巧妙的操作）
	方法1：双指针解法
	(question)，方法很巧妙
*/
func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, m+n-1
	// 从后向前遍历
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
		// nums1[k--] = nums2[j--]
		nums1[k] = nums2[j]
		k--
		j--
	}
}