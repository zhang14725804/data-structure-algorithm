import "data-structure-algorithm/algorithm/common"

/*
	集合：所有将第一个字符串的前i个字母，变成第二个字符串的前j个字母的方案
	状态计算：四种方案求最小值（insert，delete，replace（两种）），不懂
*/
func minDistance(str1 string, str2 string) int {
	n := len(str1)
	m := len(str2)
	// 二维数组（todos：我这么开辟二维数组是不是有点蠢）
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
			// 第一个字符串的第i-1个字母和第二个字符串的j-1个字母不相等
			if str1[i-1] != str2[j-1] {
				add = 1
			}
			dp[i][j] = common.Min(dp[i][j], dp[i-1][j-1]+add)
		}
	}
	return dp[n][m]
}
