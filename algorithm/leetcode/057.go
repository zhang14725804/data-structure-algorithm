/*
	插入区间

	给出一个无重叠的 ，按照区间起始端点【排序】的区间列表。
	在列表中插入一个新的区间，你需要确保列表中的区间仍然有序且不重叠（如果有必要的话，可以合并区间）。

	贪心算法（todo）：这题和他么贪心有卵子关系

	贪心算法一般用来解决需要 “找到要做某事的最小数量” 或 “找到在某些情况下适合的最大物品数量” 的问题，且提供的是无序的输入。
	贪心算法的思想是每一步都选择最佳解决方案，最终获得全局最佳的解决方案。

*/
func insert(intervals [][]int, newInterval []int) [][]int {
	newStart, newEnd := newInterval[0], newInterval[1]
	idx := 0
	n := len(intervals)
	res := make([][]int, 0)
	// intervals元素的第一个元素 小于 newInterval[0]，直接添加到结果
	for idx < n && newStart > intervals[idx][0] {
		res = append(res, intervals[idx])
		idx++
	}

	temp := make([]int, 2)
	// intervals最后一个元素的第二个元素 小于 newInterval[0]
	if len(res) == 0 || res[len(res)-1][1] < newStart {
		res = append(res, newInterval)
	} else {
		// intervals最后一个元素的第二个元素 大于 newInterval[0]，合并区间
		temp = res[len(res)-1]
		res = res[:len(res)-1]
		temp[1] = MaxInt(temp[1], newEnd)
		res = append(res, temp)
	}

	for idx < n {
		temp = intervals[idx]
		idx++
		start, end := temp[0], temp[1]
		// res最后一个元素的第二个元素 小于 start
		if res[len(res)-1][1] < start {
			res = append(res, temp)
		} else {
			// res最后一个元素的第二个元素 大于 start
			temp = res[len(res)-1]
			res = res[:len(res)-1]
			temp[1] = MaxInt(temp[1], end)
			res = append(res, temp)
		}
	}
	return res
}