/*
	todo：不懂,懵逼
	kip算法？
*/
func longestPrefix(s string) string {
	n:=len(s)
	// 增加占位符
	s= " "+s
	next:=make([]int,n+1)
	// todo：看不懂
	for i,j:=2,0;i<=n;i++{
		for j!=0 && s[i]!=s[j+1]{
			j = next[j]
		}
		if s[i]==s[j+1]{
			j++
		}
		next[i]=j
	}
	return s[1:1+next[n]]
}