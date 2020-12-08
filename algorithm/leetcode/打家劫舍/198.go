/*
	对于第 n 家店，我们只能选择偷或者不偷。

	如果偷的话，那么前 n 家商店最多能偷的钱数就是「前 n - 2 家店最多能偷的钱数」加上「第 n 家店的钱数」。因为选择偷第 n 家商店，第 n - 1 家商店就不可以偷了。

	如果不偷的话，那么前 n 家商店最多能偷的钱数就是「前 n - 1 家店最多能偷的钱数」
*/

/*
	方法1：递归。超时。递归中会计算很多重复的解
*/
var nums []int

func rob1(_nums []int) int {
	nums = _nums
	return dfs(len(nums))
}

func dfs(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}

	return MaxInt(dfs(n-2)+nums[n-1], dfs(n-1))
}

/*
	方法2：递归优化.递归中会计算很多重复的解。
	把递归过程中的解存起来，第二次遇到的话直接返回即可。

	递归：自顶向下
*/
var nums []int
var hash map[int]int // 使用数组或者hash进行cache

func rob2(_nums []int) int {
	nums = _nums
	hash = make(map[int]int, len(nums)+1)
	return dfs(len(nums))
}

func dfs(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}

	if val, ok := hahs[n]; ok {
		return val
	}

	res := MaxInt(dfs(n-2)+nums[n-1], dfs(n-1))
	hash[n] = res
	return res
}

/*
	方法3：动态规划：自底向上
	为什么有时候返回n，有时候返回n-1(question)
*/
func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}

	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = nums[0]
	// 边界条件
	for i := 2; i <= n; i++ {
		dp[i] = MaxInt(dp[i-2]+nums[i-1], dp[i-1])
	}
	return dp[n]
}

/*
	方法4：动态规划优化（状态压缩）
	更新 dp[i] 的时候，只需要 dp[i - 1] 以及 dp[i - 2] 的信息，再之前的信息就不需要了，所以我们不需要数组，只需要几个变量就可以了
*/
func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}

	pre := 0
	cur := nums[0]
	// 边界条件
	for i := 2; i <= n; i++ {
		temp := cur
		cur = MaxInt(pre+nums[i-1], cur)
		pre = temp
	}
	return cur
}