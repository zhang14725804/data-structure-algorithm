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
	// 递归出口
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
	方法2：动态规划 😅😅😅😅😅😅
	初始值 f ( 1 ) 和 f ( 2 )，然后可以求出 f ( 3 )，然后求出 f ( 4 ) ... 直到 f ( n )，一个循环就够了
*/
func climbStairs(n int) int {
	n1, n2 := 1, 2
	if n == 1 {
		return n1
	}
	if n == 2 {
		return n2
	}
	// i<=n
	for i := 3; i <= n; i++ {
		n1, n2 = n2, n1+n2
	}
	return n2
}