/*
	（空间换时间）：

	归并排序（Merge sort）是建立在归并操作上的一种有效的排序算法。该算法是采用分治法（Divide and Conquer）的一个非常典型的应用。

	分割：递归地把当前序列平均分割成两半。
	集成：在保持元素顺序的同时将上一步得到的子序列集成到一起（归并）。

	作为一种典型的分而治之思想的算法应用，归并排序的实现由两种方法：
		(1) 自上而下的递归
		(2) 自下而上的迭代；
	和选择排序一样，归并排序的性能不受输入数据的影响，但表现比选择排序好的多，因为始终都是 O(nlogn) 的时间复杂度。代价是需要额外的内存空间。
*/
func mergeSort(array []int) []int {
	length := len(array)
	if length < 2 {
		return array
	}
	key := length / 2
	// 递归左右两部份
	left := mergeSort(array[0:key])
	right := mergeSort(array[key:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	i, j := 0, 0
	m, n := len(left), len(right)
	var res []int
	for i < m && j < n {
		if left[i] > right[j] {
			res = append(res, right[j])
			j++
		} else {
			res = append(res, left[i])
			i++
		}
	}
	res = append(res, left[i:]...)
	res = append(res, right[j:]...)
	return res
}
