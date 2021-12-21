/*
	斐波那契数列
*/

/*
	方法1：迭代
	递归（自顶向下）
*/
func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	prev, cur := 1, 1
	for i := 3; i <= n; i++ {
		sum := prev + cur
		prev = cur
		cur = sum
	}
	return cur
}

/*
	方法2：带备忘录的递归（解决暴力递归带来的重叠子问题）
*/
var memo []int

func fib(N int) int {
	if N < 2 {
		return N
	}
	// 😅 question 为什么需要N+1
	memo = make([]int, N+1) // 备忘录
	return dfs(N)
}
func dfs(N int) int {
	if N == 1 || N == 2 {
		return 1
	}
	// 命中缓存
	if memo[N] != 0 {
		return memo[N]
	}
	memo[N] = dfs(N-1) + dfs(N-2)
	return memo[N]
}

/*
	方法3：动态规划（自底向上）

*/
func fib(n int) int {
	if n < 2 {
		return n
	}
	// （1）dp[i]的定义为：第i个数的斐波那契数值是dp[i]
	// 😅 数组长度n+1
	dp := make([]int, n+1)
	// （3）dp数组初始化
	dp[0] = 0
	dp[1] = 1
	// （4）确定遍历顺序。dp[i]是依赖 dp[i - 1] 和 dp[i - 2]，那么遍历的顺序一定是从前到后遍历的
	// 😅 从2开始到n，包括n
	for i := 2; i <= n; i++ {
		// （2）状态转移方程 dp[i] = dp[i - 1] + dp[i - 2];
		dp[i] = dp[i-1] + dp[i-2]
	}
	// （5）举例推导dp数组。n=10时，0 1 1 2 3 5 8 13 21 34 55
	return dp[n]
}

/*
	方法4：动态规划
	状态压缩 😅😅😅
*/
func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	// 😅😅😅 当前状态只和之前的两个状态有关，其实并不需要DP table 来存储所有的状态
	prev, cur := 1, 1
	for i := 3; i <= n; i++ {
		sum := prev + cur
		prev = cur
		cur = sum
	}
	return cur
}