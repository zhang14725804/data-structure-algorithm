/*
	给定一个区间的集合，找到需要移除区间的最小数量，使剩余区间互不重叠。

	注意:
	可以认为区间的终点总是大于它的起点。
	区间 [1,2] 和 [2,3] 的边界相互“接触”，但没有相互重叠。
*/
func intervalSchedule(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	// 按照每个区间末位值end排序
	for i := 0; i < len(intervals); i++ {
		for j := 0; j < len(intervals)-i-1; j++ {
			if intervals[j][1] > intervals[j+1][1] {
				intervals[j], intervals[j+1] = intervals[j+1], intervals[j]
			}
		}
	}
	// 至少有一个区间不相交
	count := 1
	x_end := intervals[0][1]
	for _, ints := range intervals {
		start := ints[0]
		// 如果一个区间不想与 x 的end相交，它的start必须要大于（或等于）x 的end
		if start >= x_end {
			count++
			x_end = ints[1]
		}
	}
	return count
}

func eraseOverlapIntervals(intervals [][]int) int {
	return len(intervals) - intervalSchedule(intervals)
}