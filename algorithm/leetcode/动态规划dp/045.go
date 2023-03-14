/*
	方法1:
	question 😅😅😅
*/
func jump(nums []int) int {
	ans := 0
	start := 0
	//
	end := 1
	for end < len(nums) {
		// 能跳到的最远距离
		maxPos := 0
		for i := start; i < end; i++ {
			maxPos = Max(maxPos, i+nums[i])
		}
		// 更新起跳范围
		start = end
		end = maxPos + 1
		// 跳跃次数
		ans++
	}
	return ans
}

/*
	贪婪算法：从前向后
	方法一优化版本
	question 😅😅😅
	每次在可跳范围内选择可以使得跳的更远的位置。
*/
func jump(nums []int) int {
	// end:当前能跳的边界，，
	end := 0
	// maxPos：能跳最远的距离
	maxPos := 0
	// steps：步数
	steps := 0
	for i := 0; i < len(nums)-1; i++ {
		maxPos = MaxInt(maxPos, nums[i]+i)
		// 遇到边界，更新边界，更新步数
		if i == end {
			end = maxPos
			steps++
		}
	}
	return steps
}

/*
	贪婪算法：从后向前
	最终要到达最后一个位置，然后我们找前一个位置，遍历数组，找到能到达它的位置，离它最远的就是要找的位置。
	然后继续找上上个位置，最后到了第 0 个位置就结束了
*/
func jump2(nums []int) int {
	position := len(nums) - 1
	steps := 0
	for position != 0 {
		for i := 0; i < position; i++ {
			if nums[i] >= position-i {
				position = i
				steps++
				break
			}
		}
	}
	return steps
}

/*
	自顶向下的递归动态规划
*/
var nums []int
var memo []int
var n int

func jump(_nums []int) int {
	nums = _nums
	n = len(nums)
	memo = make([]int, n)
	for i := 0; i < n; i++ {
		memo[i] = n
	}
	return dp(0)
}

func dp(p int) int {
	// base case,当p超过最后一格时，不需要跳跃
	if p >= n-1 {
		return 0
	}
	// 子问题已经计算过
	if memo[p] != n {
		return memo[p]
	}
	steps := nums[p]
	// 你可以选择跳 1 步，2 步...
	for i := 1; i <= steps; i++ {
		// 穷举每一个选择
		// 计算每一个子问题的结果
		sub := dp(p + i)
		// 取其中最小的作为最终结果
		memo[p] = MinInt(memo[p], sub+1)
	}
	return memo[p]
}