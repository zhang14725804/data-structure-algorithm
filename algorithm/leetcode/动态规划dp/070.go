/*
	假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
	每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
	注意：给定 n 是一个正整数。
*/

/*
	方法1：递归 😅😅😅😅
*/
var memo []int // 缓存 😅

func climbStairs(n int) int {
	memo = make([]int, n+1)
	return dfs(n)
}
func dfs(n int) int {
	// base case，递归出口
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	// 走一步
	n1 := 0
	if memo[n-1] == 0 {
		n1 = dfs(n - 1)
		memo[n-1] = n1
	} else {
		n1 = memo[n-1]
	}
	// 走两步
	n2 := 0
	if memo[n-2] == 0 {
		n2 = dfs(n - 2)
		memo[n-2] = n2
	} else {
		n2 = memo[n-2]
	}
	return n1 + n2
}

/*
	动态规划
*/
func climbStairs(n int) int {
	// 😅 因为下面直接对dp[2]操作了，防止空指针
	if n < 2 {
		return n
	}
	// dp[i]： 爬到第i层楼梯，有dp[i]种方法
	// 😅 数组长度n+1
	dp := make([]int, n+1)
	// dp数组如何初始化
	dp[1] = 1
	dp[2] = 2
	// 从递推公式dp[i] = dp[i - 1] + dp[i - 2];中可以看出，遍历顺序一定是从前向后遍历的
	// 😅 i从3开始
	for i := 3; i <= n; i++ {
		// 递推公式
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

/*
	方法2：动态规划
	状态压缩 😅😅😅
*/
func climbStairs(n int) int {
	prev, cur := 1, 2
	if n == 1 {
		return prev
	}
	if n == 2 {
		return cur
	}
	// i<=n
	for i := 3; i <= n; i++ {
		prev, cur = cur, prev+cur
	}
	return cur
}

/*
	拓展（完全背包问题 😅）：一步一个台阶，两个台阶，三个台阶，直到 m个台阶，有多少种方法爬到n阶楼顶

	class Solution {
	public:
		int climbStairs(int n) {
			vector<int> dp(n + 1, 0);
			dp[0] = 1;
			for (int i = 1; i <= n; i++) {
				for (int j = 1; j <= m; j++) { // 把m换成2，就可以AC爬楼梯这道题
					if (i - j >= 0) dp[i] += dp[i - j];
				}
			}
			return dp[n];
		}
	};
*/
