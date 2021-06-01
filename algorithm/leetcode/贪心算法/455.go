/*
	这里的局部最优就是大饼干喂给胃口大的，充分利用饼干尺寸喂饱一个，全局最优就是喂饱尽可能多的小孩。

	可以尝试使用贪心策略，先将饼干数组和小孩数组排序。
	然后从后向前遍历小孩数组，用大饼干优先满足胃口大的，并统计满足小孩数量。
*/
func findContentChildren(g []int, s []int) int {
	g = BubbleSort(g)
	s = BubbleSort(s)
	// 饼干数组下标
	sLen := len(s) - 1
	gLen := len(g) - 1
	var res int
	for i := gLen; i >= 0; i-- {
		if sLen >= 0 && s[sLen] >= g[i] {
			// 😅 遍历饼干并没有再起一个for循环，而是采用自减的方式
			sLen--
			res++
		}
	}
	return res
}

// 小饼干先喂饱小胃口
func findContentChildren(g []int, s []int) int {
	g = BubbleSort(g)
	s = BubbleSort(s)
	index := 0
	for i := 0; i < len(s); i++ {
		if index < len(g) && g[index] <= s[i] {
			index++
		}
	}
	return index
}