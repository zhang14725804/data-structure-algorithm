func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	// 按照每个区间末位值end排序
	for i := 0; i < len(points); i++ {
		for j := 0; j < len(points)-i-1; j++ {
			if points[j][1] > points[j+1][1] {
				points[j], points[j+1] = points[j+1], points[j]
			}
		}
	}
	// 至少有一个区间不相交
	count := 1
	x_end := points[0][1]
	for _, ints := range points {
		start := ints[0]
		// 如果一个区间不想与 x 的end相交，它的start必须要大于x 的end
		if start > x_end {
			count++
			x_end = ints[1]
		}
	}
	return count
}