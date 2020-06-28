/*
	给定一个字符串 s，将 s 分割成一些子串，使每个子串都是回文串。
	返回 s 所有可能的分割方案。

	暴力搜索（方案数是指数级别）
	暴力搜索需要考虑的（😅）
	（1）状态
	（2）当前正在枚举的区间
	（3）枚举到的字符下标

	有问题：
	todo：aab，结果是：[["aa","b","b"],["aa","b"]]
*/
var path []string
var ans [][]string
func partition(s string) [][]string {
	dfs("", 0, s)
	return ans
}
// 当前区间是否是回文串
func check(now string) bool {
	if len(now) == 0{
		return false
	}
	for i,j:=0,len(now)-1; i<j; i,j=i+1,j-1{
		if now[i] != now[j]{
			return false
		}
	}
	return true
}
// 当前方案，当前位置，字符串
func dfs(now string, u int, s string)  {
	// 已经枚举完了
	if u == len(s){
		if check(now) == true{
			path = append(path,now)
			ans = append(ans,path)
			path = path[:len(path)-1]
		}
		return 
	}
	// 
	if check(now) == true{
		path = append(path,now)
		dfs("",u,s)
		path = path[:len(path)-1]
	}
	dfs(now+string(s[u]), u+1, s)
}