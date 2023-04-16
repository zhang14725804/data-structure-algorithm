/*
	方法1: 【自顶向下】的递归
	1. dfs：返回 s1[0..i] 和 s2[0..j] 的最小编辑距离
*/
func minDistance(s1 string, s2 string) int {

	var dfs func(i int, j int) int
	dfs = func(i int, j int) int {
		// i,j从0开始，所以这里要【+1】操作
		if i < 0 {
			return j + 1
		}
		if j < 0 {
			return i + 1
		}

		if s1[i] == s2[j] {
			return dfs(i-1, j-1)
		} else {
			insert := dfs(i, j-1) + 1    // 插入
			delete := dfs(i-1, j) + 1    // 删除
			replace := dfs(i-1, j-1) + 1 // 替换
			return min(min(insert, delete), replace)
		}
	}
	return dfs(len(s1)-1, len(s2)-1)
}

/*
	方法2：带【cache】的递归
	有问题，应该是缓存的hash key有问题
*/
func minDistance(s1 string, s2 string) int {

	var dfs func(i int, j int) int
	cache := map[string]int{}
	dfs = func(i int, j int) int {
		if i < 0 {
			return j + 1
		}
		if j < 0 {
			return i + 1
		}

		key := fmt.Sprintf("%d_%d", i, j)
		if dist, ok := cache[key]; ok {
			return dist
		} else if s1[i] == s2[j] {
			cache[key] = dfs(i-1, j-1)
		} else {
			insert := dfs(i, j-1) + 1
			delete := dfs(i-1, j) + 1
			replace := dfs(i-1, j-1) + 1
			cache[key] = min(min(insert, delete), replace)
		}
		return cache[key]
	}
	return dfs(len(s1)-1, len(s2)-1)
}

/*
	方法3:动态规划
	【自底向上】
	1. 状态
	2. 选择
	3. 状态转移方程
	4. base case
*/
func minDistance(s1 string, s2 string) int {
	m, n := len(s1), len(s2)
	// dp[i][j] 存储 s1[0..i-1] 变成 s2[0..j-1] 的最小编辑距离
	dp := make([][]int, m+1)
	for i, _ := range dp {
		dp[i] = make([]int, n+1)
	}

	// base case
	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}

	//
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				insert := dp[i][j-1] + 1    // 插入
				delete := dp[i-1][j] + 1    // 删除
				replace := dp[i-1][j-1] + 1 // 替换
				dp[i][j] = min(min(insert, delete), replace)
			}
		}
	}

	//
	return dp[m][n]
}