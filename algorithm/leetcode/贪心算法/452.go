/*
局部最优：当气球出现重叠，一起射，所用弓箭最少。全局最优：把所有气球射爆所用弓箭最少。
*/

func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	points = BubbleSort(points)
	// points 不为空至少需要一支箭
	res := 1
	for i := 1; i < len(points); i++ {
		// 气球i和气球i-1不挨着，注意这里不是>=
		if points[i][0] > points[i-1][1] {
			res++
		} else {
			// 气球i和气球i-1挨着
			// 更新重叠气球最小右边界。如果气球重叠了，重叠气球中右边边界的最小值 之前的区间一定需要一个弓箭。
			points[i][1] = MinInt(points[i-1][1], points[i][1])
		}
	}
	return res
}

// 按照起始位置排序，那么就从前向后遍历气球数组，靠左尽可能让气球重复。
func compare(a, b []int) bool {
	return a[0] < b[0]
}