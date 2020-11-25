/*
	方法2：动态规划
	集合：所有将第一个字符串的前i个字母，变成第二个字符串的前j个字母的方案
	状态计算：三种方案求最小值
		最后一步insert，dp[i][j-1] + 1
		最后一步delete，dp[i-1][j] + 1
		最后一步replace（两种，str1[i]和str2[j]已经相等了，str1[i]和str2[j]不相等）），有点难理解

*/
func minDistance(str1 string, str2 string) int {
	n := len(str1)
	m := len(str2)
	// (question)为什么要开辟n+1,m+1的空间
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, m+1)
	}
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
			dp[i][j] = common.Min(dp[i-1][j], dp[i][j-1]) + 1
			// replace
			var add int
			// 第一个字符串的第i-1个字母和第二个字符串的j-1个字母不相等（两个字符串的最后一个字母不相等）
			if str1[i-1] != str2[j-1] {
				add = 1
			}
			dp[i][j] = common.Min(dp[i][j], dp[i-1][j-1]+add)
		}
	}
	return dp[n][m]
}

/*
	方法1：递归
	超出时间限制了😅
*/
func minDistance(word1 string, word2 string) int {
	// 三种边界条件，无需赘述边界条件
	if len(word1) == 0 && len(word2) == 0 {
		return 0
	}
	if len(word1) == 0 {
		return len(word2)
	}
	if len(word2) == 0 {
		return len(word1)
	}
	// 这他么怎么理解
	x := minDistance(word1, word2[:len(word2)-1]) + 1
	y := minDistance(word1[:len(word1)-1], word2) + 1
	z := minDistance(word1[:len(word1)-1], word2[:len(word2)-1])
	if word1[len(word1)-1] != word2[len(word2)-1] {
		z++
	}
	return compare(compare(x, y, false), z, false)
}
