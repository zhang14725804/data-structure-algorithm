/*
	方法2：动态规划
	集合：所有将第一个字符串的前i个字母，变成第二个字符串的前j个字母的方案
	状态计算：三种方案求最小值
		最后一步insert，dp[i][j-1] + 1
		最后一步delete，dp[i-1][j] + 1
		最后一步replace（两种，str1[i]和str2[j]已经相等了，str1[i]和str2[j]不相等）），有点难理解
	0109 未做
	todo：
		（1）状态压缩
		（2）这里只求出了最小的编辑距离，那具体的操作是什么
*/
func minDistance(str1 string, str2 string) int {
	n := len(str1)
	m := len(str2)
	// (question)为什么要开辟n+1,m+1的空间
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, m+1)
	}
	// base case
	// 第一个字符串前i个字母变成第二个字符串的前0个字母
	for i := 0; i <= n; i++ {
		dp[i][0] = i
	}
	// 第一个字符串前0个字母变成第二个字符串的前i个字母
	for i := 0; i <= m; i++ {
		dp[0][i] = i
	}
	// i和j从1开始
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			// insert和delete
			dp[i][j] = MinInt(dp[i-1][j], dp[i][j-1]) + 1
			// replace
			var add int
			// 第一个字符串的第i-1个字母和第二个字符串的j-1个字母不相等（两个字符串的最后一个字母不相等）
			if str1[i-1] != str2[j-1] {
				add = 1
			}
			dp[i][j] = MinInt(dp[i][j], dp[i-1][j-1]+add)
		}
	}
	return dp[n][m]
}

/*
	方法1：递归（自顶向下），存在重叠子问题
	超出时间限制了😅
	minDistance：返回 s1[0..i] 和 s2[0..j] 的最小编辑距离
*/
func minDistance1(word1 string, word2 string) int {
	// 都是空串、S1或S2为空
	if len(word1) == 0 && len(word2) == 0 {
		return 0
	}
	if len(word1) == 0 {
		return len(word2)
	}
	if len(word2) == 0 {
		return len(word1)
	}
	// 插入
	x := minDistance(word1, word2[:len(word2)-1]) + 1
	// 替换
	y := minDistance(word1[:len(word1)-1], word2) + 1
	// word1[i]和word2[j]相同，什么都不需要做
	z := minDistance(word1[:len(word1)-1], word2[:len(word2)-1])
	// word1[i]和word2[j]不相同
	if word1[len(word1)-1] != word2[len(word2)-1] {
		z++
	}
	return MinInt(MinInt(x, y), z)
}

/*
	方法2：带cache的递归
	有问题，应该是缓存的hash key有问题
*/
func minDistance2(s1 string, s2 string) int {
	var dfs func(i, j int) int
	// 缓存
	hash := make(map[string]int)
	dfs = func(i, j int) int {
		if j == -1 {
			return i + 1
		}
		if i == -1 {
			return j + 1
		}
		if val, ok := hash[fmt.Sprintf("%v", i)+fmt.Sprintf("%v", j)]; ok {
			return val
		}
		if s1[i] == s2[j] {
			hash[fmt.Sprintf("%v", i)+fmt.Sprintf("%v", j)] = dfs(i-1, j-1)
		} else {
			hash[fmt.Sprintf("%v", i)+fmt.Sprintf("%v", j)] = MinInt(MinInt(dfs(i-1, j)+1, dfs(i, j-1)+1), dfs(i-1, j-1)+1)
		}
		return hash[fmt.Sprintf("%v", i)+fmt.Sprintf("%v", j)]
	}
	// i，j 初始化指向最后一个索引
	return dfs(len(s1)-1, len(s2)-1)
}