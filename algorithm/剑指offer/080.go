/*
	将一个骰子投掷n次，获得的总点数为s，s的可能范围为n~6n。

	掷出某一点数，可能有多种掷法，例如投掷2次，掷出3点，共有[1,2],[2,1]两种掷法。

	请求出投掷n次，掷出n~6n点分别有多少种掷法。
	todo：难理解（背包问题,分组背包问题😅）
	两种思路：递归（dfs）、动态规划（dp）
	递归（dfs）步骤：状态，计算状态
	动态规划（dp）步骤：状态，如何计算状态，边界
*/
func numberOfDice(n int) []int {
	f:=make([][]int,n+1)
	for i := 0; i < n+1; i++ {
		f[i] = make([]int,n*6+1)
	}

	f[0][0] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= i*6; j++ {
			for k := 1; k <= min(j,6); k++ {
				f[i][j] += f[i-1][j-k]
			}
		}
	}
	res := make([]int,0)
	for i := n; i <= n*6; i++ {
		res = append(res,f[n][i])
	}
	return res
}

func min(a,b int) int{
	if a>b{
		return b
	}
	return a
}
