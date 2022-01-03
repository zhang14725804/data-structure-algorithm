/*
	42. Trapping Rain Water（84题的进阶版）
	给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

	思路1：按行求（会超时）
	思路2：按列求
	思路3：动态规划
	思路4：双指针
	思路5：单调栈

	0103 先跳过
*/

/*
	方法1：暴力破解
	对于这种问题，我们不要想整体，而应该去想局部；
	就像之前的文章写的动态规划问题处理字符串问题，不要考虑如何处理整个字符串，而是去思考应该如何处理每一个字符。
	对于位置i，能够装的水为：
	water[i] = min(
		# 左边最高的柱子
		max(height[0..i]),
		# 右边最高的柱子
		max(height[i..end])
	) - height[i]
	时间复杂度 O(N^2)，空间复杂度 O(1)
*/
func trap1(height []int) int {
	n := len(height)
	res := 0
	// 边界条件
	for i := 1; i < n-1; i++ {
		lMax, rMax := 0, 0
		// 边界条件
		for j := i; j < n; j++ {
			rMax = MaxInt(height[j], rMax)
		}
		// 边界条件
		for k := i; k >= 0; k-- {
			lMax = MaxInt(height[k], lMax)
		}
		res += MinInt(lMax, rMax) - height[i]
	}
	return res
}

/*
	方法2：在暴力破解的基础之上，增加cache
	时间复杂度O(N)，已经是最优了，但是空间复杂度是 O(N)
*/
func trap2(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}
	res := 0
	lMax, rMax := make([]int, n), make([]int, n)
	// l_max[i]和r_max[i]分别代表height[0..i]和height[i..n-1]的最高柱子高度
	lMax[0] = height[0]
	rMax[n-1] = height[n-1]
	// 边界条件
	for i := 1; i < n; i++ {
		lMax[i] = MaxInt(height[i], lMax[i-1])
	}
	// 边界条件
	for k := n - 2; k >= 0; k-- {
		rMax[k] = MaxInt(height[k], rMax[k+1])
	}
	// 边界条件
	for i := 1; i < n; i++ {
		res += MinInt(lMax[i], rMax[i]) - height[i]
	}
	return res
}

/*
	方法3：双指针算法
	question
*/
func trap(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}
	res := 0
	left, right := 0, n-1
	// l_max和r_max代表的是height[0..left]和height[right..n-1]的最高柱子高度
	// 只在乎min(l_max, r_max)。对于上图的情况，我们已经知道l_max < r_max了，至于这个r_max是不是右边最大的，不重要。
	// 重要的是height[i]能够装的水只和较低的l_max之差有关
	lMax, rMax := height[0], height[n-1]
	for left < right {
		lMax = MaxInt(lMax, height[left])
		rMax = MaxInt(rMax, height[right])
		if lMax < rMax {
			res += lMax - height[left]
			left++
		} else {
			res += rMax - height[right]
			right--
		}
	}

	return res
}

/*
	question 单调栈？
*/
func trap_question(height []int) int {
	res := 0
	stk := &stack{}

	for i := 0; i < len(height); i++ {
		last := 0
		for stk.size() > 0 && height[stk.top()] < height[i] {
			t := stk.top()
			stk.pop()
			res += (i - t - 1) * (height[t] - last)
			last = height[t]
		}
		if stk.size() > 0 {
			res += (i - stk.top() - 1) * (height[i] - last)
		}
		stk.push(i)
	}
	return res
}

