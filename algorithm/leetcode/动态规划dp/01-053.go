/*
	方法1：暴力枚举
	数量太大就不行了😅
*/
func maxSubArray1(nums []int) int {
	res := INT_MIN
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			// 计算i~j元素的和
			sum := 0
			for k := i; k <= j; k++ {
				sum += nums[k]
			}
			res = compare(res, sum, true)
		}
	}
	return res
}

/*
	方法2：动态规划
	（1）状态表示，包括集合（所有以i结尾的子段）、属性（Max）
	（2）状态计算 f[i] = max(f[i-1], 0) + nums[i]（f[i-1]取或者不取的最大值，再加上nums[i]）
*/
func maxSubArray(nums []int) int {
	nLen := len(nums)
	if nLen == 0 {
		return 0
	}

	// 以 nums[i] 为结尾的「最大子数组和」为 dp[i]。
	dp := make([]int, nLen)
	// base case
	dp[0] = nums[0]
	// 初始值
	res := dp[0]
	for i := 1; i < nLen; i++ {
		dp[i] = Max(nums[i], nums[i]+dp[i-1])
		// 更新结果
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}

func maxSubArray(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	dp := make([]int, n)
	dp[0] = nums[0]
	res := dp[0]
	// 😅 一个循环搞定
	for i := 1; i < n; i++ {
		dp[i] = MaxInt(nums[i], nums[i]+dp[i-1])
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}

/*
	方法3：动态规划 + 状态压缩（😅😅😅）
	ps:dp[i] 仅仅和 dp[i-1] 的状态有关
*/
func maxSubArray3(nums []int) int {
	res := INT_MIN
	last := 0
	for i := 0; i < len(nums); i++ {
		now := MaxInt(last, 0) + nums[i]
		res = MaxInt(res, now)
		last = now
	}
	return res
}

/*
	方法1：贪心算法

	如果 -2 1 在一起，计算起点的时候，一定是从1开始计算，因为负数只会拉低总和，这就是贪心贪的地方！

	局部最优：当前“连续和”为负数的时候立刻放弃，从下一个元素重新计算“连续和”，因为负数加上下一个元素 “连续和”只会越来越小。
	全局最优：选取最大“连续和”
	局部最优的情况下，并记录最大的“连续和”，可以推出全局最优。

	从代码角度上来讲：遍历nums，从头开始用count累积，如果count一旦加上nums[i]变为负数，那么就应该从nums[i+1]开始从0累积count了，因为已经变为负数的count，只会拖累总和。
	这相当于是暴力解法中的不断调整最大子序和区间的起始位置。
	那有同学问了，区间终止位置不用调整么？ 如何才能得到最大“连续和”呢？
	区间的终止位置，其实就是如果count取到最大值了，及时记录下来了。
*/
func maxSubArray(nums []int) int {
	res := -(1 << 32)
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		// 取区间累计的最大值（相当于不断确定最大子序终止位置）
		if res < sum {
			res = sum
		}
		// 相当于重置最大子序起始位置，因为遇到负数一定是拉低总和
		if sum <= 0 {
			sum = 0
		}
	}
	return res
}

/*
	方法二：动态规划（超出时间限制）
	用一个二维数组 dp[ i ] [ len ] 表示从下标 i 开始，长度为 len 的子数组的元素和。
	这样长度是 len + 1 的子数组就可以通过长度是 len 的子数组去求，也就是下边的递推式，
	dp [ i ] [ len + 1 ] = dp[ i ] [ len ] + nums [ i + len - 1 ]。
	考虑到求 i + 1 的情况的时候，我们只需要 i 时候的情况，所有我们其实没必要用一个二维数组，直接用一维数组就可以了
	0102 我居然用双指针（愚蠢）
	没懂
*/

func maxSubArray(nums []int) int {
	n := len(nums)
	res := -(1 << 32)
	dp := make([]int, n)
	// 遍历长度
	for i := 1; i <= n; i++ {
		// 遍历下标
		for j := 0; j <= n-i; j++ {
			dp[j] = dp[j] + nums[i+j-1]
			if res < dp[j] {
				res = dp[j]
			}
		}
	}
	return res
}