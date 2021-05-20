/*
	斐波那契数列

	方法：动态规划（自底向上）、递归（自顶向下）
*/

/*
	方法1：迭代
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
	// question 为什么需要N+1
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
func fib(N int) int {
	if N < 2 {
		return N
	}
	dp := make([]int, N+1)
	dp[1], dp[2] = 1, 1
	// n从3开始，n+1个而不是n个 😅😅😅
	for i := 3; i <= N; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[N]
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
	// 根据斐波那契数列的状态转移方程，当前状态只和之前的两个状态有关，其实并不需要DP table 来存储所有的状态
	prev, cur := 1, 1
	for i := 3; i <= n; i++ {
		sum := prev + cur
		prev = cur
		cur = sum
	}
	return cur
}