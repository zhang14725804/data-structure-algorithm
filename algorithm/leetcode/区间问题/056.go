/*
	给出一个区间的集合，请合并所有重叠的区间。每个数组只有两个数

	思路：
	将所有数组按照数组第一个数字大小排序
	定义curr，代表当前的数组，初始值是intervals[0]
	如果curr[1] >= intervals[i][0],触发合并条件
*/
func merge(intervals [][]int) [][]int {
	if len(intervals) < 2 {
		return intervals
	}
	// （1） 根据区间【左】边界排序
	for i := 0; i < len(intervals)-1; i++ {
		for j := 0; j < len(intervals)-i-1; j++ {
			if intervals[j][0] > intervals[j+1][0] {
				intervals[j], intervals[j+1] = intervals[j+1], intervals[j]
			}
		}
	}
	// （2） 当前数组
	curr := intervals[0]
	res := make([][]int, 0)

	// （3） 找出相交的情况
	for i := 1; i < len(intervals); i++ {
		// 相交的情况。 【右】边界 大于等于 另一个数组的第一个数；右边取二者的最大值
		if curr[1] >= intervals[i][0] {
			curr[1] = MaxInt(curr[1], intervals[i][1])
		} else {
			// 没有相交
			res = append(res, curr)
			curr = intervals[i]
		}
	}
	// （4） 处理剩余的数组 😅😅
	if len(curr) != 0 {
		res = append(res, curr)
	}

	return res
}
