/*
	思路不错
*/
var ans []string
func letterCasePermutation(S string) []string {
	dfs(S,0)
	return ans
}
func dfs(S string,u int)  {
	if u == len(S){
		ans = append(ans,S)
		return
	}
	// 不变的情况递归一次
	dfs(S,u+1)
	// 变大写之后再次递归
	if S[u] >= 'A'{
		// 改大小写
		S = updateString(S,u)
		dfs(S,u+1)
	}
}

// 修改字符串
func updateString(S string,u int) string {
	c := []byte(S)  // 将字符串 s 转换为 []byte 类型
    c[u] ^= 32
    return string(c)  // 再转换回 string 类型
}