/*
	归并排序【空间换时间】基本思想：
		归并排序是建立在归并操作上的一种有效的排序算法。
		该算法是采用【分治法（Divide and Conquer）】的一个非常典型的应用。
		将已有序的子序列合并，得到完全有序的序列；即先使每个子序列有序，再使子序列段间有序。若将两个有序表合并成一个有序表，称为2-路归并。

	分割：递归地把当前序列平均分割成两半。
	集成：在保持元素顺序的同时将上一步得到的子序列集成到一起（归并）。

	作为一种典型的分而治之思想的算法应用，归并排序的实现由两种方法组成：
		(1) 自上而下的递归
		(2) 自下而上的迭代；
*/
func mergeSort(nums []int) []int {
	n := len(nums)
	// 😅base case，递归出口
	if n < 2 {
		return nums
	}
	// 😅 选择分界点
	key := n / 2
	// 😅 递归左右两部份
	left := mergeSort(nums[0:key])
	right := mergeSort(nums[key:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	i, j := 0, 0
	m, n := len(left), len(right)
	var res []int
	// 从左右两个数组中根据大小每次选择一个元素
	for i < m && j < n {
		if left[i] > right[j] {
			res = append(res, right[j])
			j++
		} else {
			res = append(res, left[i])
			i++
		}
	}
	// 😅😅😅 拼接剩余的元素
	res = append(res, left[i:]...)
	res = append(res, right[j:]...)
	return res
}
