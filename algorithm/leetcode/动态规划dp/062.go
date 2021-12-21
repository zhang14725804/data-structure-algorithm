/*
	一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
	机器人每次只能向下或者向右移动一步。 机器人试图达到网格的右下角（在下图中标记为“Finish”）。
	问总共有多少条不同的路径？
*/

/*
	方法2：自底向上的递归 😅😅😅😅😅😅

	求 ( 0 , 0 ) 点到（ m - 1 , n - 1） 点的走法。

	（0，0）点到（m - 1 , n - 1） 点的走法等于（0，0）点【右边】的点 （1，0）到（m - 1 , n - 1）的走法加上（0，0）点【下边】的点（0，1）到（m - 1 , n - 1）的走法。

	而左边的点（1，0）点到（m - 1 , n - 1） 点的走法等于（2，0） 点到（m - 1 , n - 1）的走法加上（1，1）点到（m - 1 , n - 1）的走法。

	下边的点（0，1）点到（m - 1 , n - 1） 点的走法等于（1，1）点到（m - 1 , n - 1）的走法加上（0，2）点到（m - 1 , n - 1）的走法。

	然后一直递归下去，直到 （m - 1 , n - 1） 点到（m - 1 , n - 1） ，返回 1。
*/
var visited map[string]int // 缓存已经走过的路径
var m, n int

func uniquePaths(_m int, _n int) int {
	visited = make(map[string]int)
	m, n = _m, _n
	return dfs(0, 0)
}

func dfs(x, y int) int {
	// 😅 base case 递归出口
	if x == m-1 && y == n-1 {
		return 1
	}
	// 向左、向下的走法
	right, bottom := 0, 0
	key := fmt.Sprintf("%v", x+1) + "@" + fmt.Sprintf("%v", y)
	// 😅 向右探索所有结果
	if val, ok := visited[key]; !ok {
		// 判断边界
		if x+1 <= m {
			right = dfs(x+1, y)
		}
	} else {
		right = val
	}

	key = fmt.Sprintf("%v", x) + "@" + fmt.Sprintf("%v", y+1)
	// 😅 向下探索所有结果
	if val, ok := visited[key]; !ok {
		// 判断边界
		if y+1 <= n {
			bottom = dfs(x, y+1)
		}
	} else {
		bottom = val
	}

	key = fmt.Sprintf("%v", x) + "@" + fmt.Sprintf("%v", y)
	visited[key] = right + bottom
	return bottom + right
}

/*
	方法1：动态规划
*/
func uniquePaths(m int, n int) int {
	// 😅😅😅 dp[i][j] ：表示从（0 ，0）出发，到(i, j) 有dp[i][j]条不同的路径。
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	// 😄 dp数组的初始化 首先dp[i][0]一定都是1，因为从(0, 0)的位置到(i, 0)的路径只有一条，那么dp[0][j]也同理
	for row := 0; row < m; row++ {
		dp[row][0] = 1
	}
	for col := 0; col < n; col++ {
		dp[0][col] = 1
	}
	// 遍历顺序：从左到右从上到下
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			// 想要求dp[i][j]，只能有两个方向来推导出来，即dp[i - 1][j] 和 dp[i][j - 1]
			// 😅 我想成了 dp[i][j] = dp[i+1][j] + dp[i][j+1]
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

/*
	方法3，动态规划
	😄😄😄 question 状态压缩
*/
func uniquePaths(m int, n int) int {
	dp := make([]int, m)
	// 初始化最后一列
	for i := 0; i < m; i++ {
		dp[i] = 1
	}
	// 😅😅😅 从右向左更新所有列 或者 从左向右都可以
	for i := n - 2; i >= 0; i-- {
		for j := m - 2; j >= 0; j-- {
			dp[j] = dp[j] + dp[j+1]
		}
	}
	return dp[0]
}