/*
	删除被覆盖区间
	给你一个区间列表，请你删除列表中被其他区间所覆盖的区间。
	只有当 c <= a 且 b <= d 时，我们才认为区间 [a,b) 被区间 [c,d) 覆盖。
	在完成所有删除操作后，请你返回列表中剩余区间的数目。

	1 <= intervals.length <= 1000
	0 <= intervals[i][0] < intervals[i][1] <= 10^5
	对于所有的 i != j：intervals[i] != intervals[j]
*/

/*
	技巧：按照起点升序排序，若起点相同，按照终点降序排序
*/
func removeCoveredIntervals1(intervals [][]int) int {
	n := len(intervals)
	// 起点升序，若相同，终点降序
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			// 起点不同
			if intervals[j][0] > intervals[j+1][0] {
				intervals[j], intervals[j+1] = intervals[j+1], intervals[j]
			} else if intervals[j][0] == intervals[j+1][0] {
				// 起点相同，按照终点降序排序
				if intervals[j+1][1] > intervals[j][1] {
					intervals[j], intervals[j+1] = intervals[j+1], intervals[j]
				}
			}
		}
	}
	// 注意循环条件len(intervals)-1，而不是n-1，因为数组是变化的
	for i := 0; i < len(intervals)-1; i++ {
		// 有包含的区间，删除对应数组，更新数组
		if intervals[i][1] >= intervals[i+1][1] {
			intervals = append(intervals[:i+1], intervals[i+2:]...)
			i--
		}
	}
	return len(intervals)
}

func removeCoveredIntervals(intervals [][]int) int {
	n := len(intervals)
	// 起点升序，若相同，终点降序
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			// 起点不同
			if intervals[j][0] > intervals[j+1][0] {
				intervals[j], intervals[j+1] = intervals[j+1], intervals[j]
			} else if intervals[j][0] == intervals[j+1][0] {
				// 起点相同，按照终点降序排序
				if intervals[j+1][1] > intervals[j][1] {
					intervals[j], intervals[j+1] = intervals[j+1], intervals[j]
				}
			}
		}
	}

	// 记录合并区间的起点和终点
	left, right := intervals[0][0], intervals[0][1]
	res := 0

	for i := 1; i < n; i++ {
		intv := intervals[i]
		// 找到覆盖区间
		if left <= intv[0] && right >= intv[1] {
			res++
		}
		// 找到交叉区间
		if right >= intv[0] && right <= intv[1] {
			right = intv[1]
		}
		// 完全不相交，更新起点终点
		if right < intv[0] {
			left, right = intv[0], intv[1]
		}
	}
	return n - res
}