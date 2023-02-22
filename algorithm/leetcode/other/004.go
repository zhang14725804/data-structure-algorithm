/*
	给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。
	请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。
*/
// 解法1：先将两个数组合并，两个有序数组的合并也是归并排序中的一部分。然后根据奇数，还是偶数，返回中位数。
func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	n, m := len(nums1), len(nums2)
	// nums1或nums2为空的情况
	if m == 0 {
		if n%2 == 0 {
			return float64((nums1[n/2] + nums1[n/2-1])) / 2.0
		} else {
			return float64(nums1[n/2])
		}
	}
	if n == 0 {
		if m%2 == 0 {
			return float64((nums2[m/2] + nums2[m/2-1])) / 2.0
		} else {
			return float64(nums2[m/2])
		}
	}

	// 归并排序
	nums := make([]int, n+m)
	// count目前已经排序元素数量，i、j分别代表当前遍历到nums1，nums2元素下标
	count, i, j := 0, 0, 0
	for count < (m + n) {
		// 已经遍历完nums1或nums2
		if i == n {
			for j < m {
				nums[count] = nums2[j]
				count++
				j++
			}
			continue
		}
		if j == m {
			for i < n {
				nums[count] = nums1[i]
				count++
				i++
			}
			continue
		}
		// 合并nums1，nums2的数据

		if nums1[i] < nums2[j] {
			nums[count] = nums1[i]
			i++
		} else {
			nums[count] = nums2[j]
			j++
		}
		count++
	}
	if count%2 == 0 {
		return float64((nums[count/2] + nums[count/2-1])) / 2.0
	} else {
		return float64(nums[count/2])
	}
}

// 解法2：二分法（todo）
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

}
