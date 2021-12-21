/*
	总结如下难点：
	难点一：一看题就有感觉需要排序，但究竟怎么排序，按左边界排还是右边界排。
	难点二：排完序之后如何遍历，如果没有分析好遍历顺序，那么排序就没有意义了。
	难点三：直接求重复的区间是复杂的，转而求最大非重复区间个数。
	难点四：求最大非重复区间个数时，需要一个分割点来做标记。

	按照右边界排序，从左向右记录非交叉区间的个数。最后用区间总数减去非交叉区间的个数就是需要移除的区间个数了。

	右边界排序之后，局部最优：优先选右边界小的区间，所以从左向右遍历，留给下一个区间的空间大一些，从而尽量避免交叉。全局最优：选取最多的非交叉区间。

	每次取非交叉区间的时候，都是可右边界最小的来做分割点（这样留给下一个区间的空间就越大）
*/

func eraseOverlapIntervals(intervals [][]int) int {
	iLen := len(intervals)
	if iLen == 0 {
		return 0
	}
	intervals = BubbleSort(intervals)
	// 记录非交叉区间的个数
	count := 1
	// 记录区间分割点
	end := intervals[0][1]
	for i := 1; i < iLen; i++ {
		if end <= intervals[i][0] {
			count++
			end = intervals[i][1]
		}
	}
	return iLen - count
}

func compare(a, b []int) bool {
	return a[1] < b[1]
}