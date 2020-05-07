/*
	把数字翻译成字符串
	todo：又是动态规划问题😅
	（1）状态标识
	（2）状态如何计算
	（3）边界

	todo：难的很，很难
*/
func getTranslationCount(s string) int {
	n := len(s)
	f := make([]int,n+1)
	f[0] = 1
	for i:=1;i<=n;i++{
		f[i] = f[i-1]
		// 把字符串变成数字
		if i>1{
			t := (s[i-2]-'0')*10 + s[i-1]-'0'
			if t>=10 && t <=25{
				f[i] += f[i-2]
			}
		}
		
	}
	return f[n]
}