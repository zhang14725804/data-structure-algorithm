/*
	给定一个字符串 s，将 s 分割成一些子串，使每个子串都是回文串。
	返回符合要求的最少分割次数。

	难，动态规划的思路
	todo：思路不懂，代码也有问题
*/
func minCut(s string) int {
    s = " " + s
	n:=len(s)
	f := make([]int,n+1)
	for i:=0;i<n;i++{
		f[i] = INT_MAX
	}
	dp := make([][]bool,n+1)
	for i:=0;i<n+1;i++{
		dp[i] = make([]bool,n+1)
	}
	// 考虑所有长度的子串（最小长度从1开始）
	for len:=1;len<=n;len++{
		// 从每个下标开始
		for i:=0;i<=n-len;i++{
			j:=i+len-1
			// i + 1 <= j - 1 并且 j:=i+len-1，所有len<3
			dp[i][j] = s[i]==s[j] && (len < 3 || dp[i+1][j-1])
		}
	}

	f[0] = 0
	for i:=1;i<=n;i++{
		for j:=1;j<=i;j++{
			if dp[j][i]{
				f[i] = MinInt(f[i],f[j-1]+1)
			}
		}
	}
	return f[n]-1
}
