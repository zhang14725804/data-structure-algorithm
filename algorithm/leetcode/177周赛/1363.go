/*
	todo：背包问题求具体问题。dp动态规划问题：实质是暴力枚举所有可能😅。这个要多看看！
	这个题不懂。代码存在问题
*/
func largestMultipleOfThree(digits []int) string {
	n:=len(digits)
	sort.Ints(digits)
	f:=make([][]int,n+1)
	for i := 0; i < n+1; i++ {
		f[i] = make([]int, 3)
	}
	f[0][1] , f[0][2] = 999999,999999
	// 
	for i:=1;i<=n;i++{
		for j:=0;j<3;j++{
			f[i][j] = int(math.Max(float64(f[i-1][j]),float64(f[i-1][(j-digits[i])%3]+1)))
		}
	}
	if f[n][0]<=0{
		return ""
	}
	var res string
	for i,j:=n,0;i>0;i--{
		if f[i-1][j] ==f[i-1][(j-digits[i])%3]+1{
			res+=string(digits[i-1])
			j = (j-digits[i])%3
			if res=="0"{
				return res
			}
		}
	}
	return res
}
