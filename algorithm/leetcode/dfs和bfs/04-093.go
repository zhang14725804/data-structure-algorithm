/*
	todo 这个题。。。不懂
	代码有问题，无非就是golang中byte，rune和string之间的破事
*/ 
var ans []string
func restoreIpAddresses(s string) []string {
	var path string 
	dfs(s,0,0,path)
	return ans
}

// 
func dfs(s string, u,k int, path string)  {
	if u == len(s){
		if k == 4{
			ans = append(ans,path[1:])
			return
		}
	}
	if k>4{
		return
	}
	// 当前这一位是0，只能是单独的一个数，例如：103.0.169.39
	if s[u]=='0'{
		// 
		dfs(s, u+1, k+1, path+".0")
	}else{
		// 23，1；23*10 + 1
		for i,t := u,0; i < len(s); i++ {
			// s[i]-'0' 转数数字 这里有问题
			t = t*10 + s[i]-'0'
			if t<256{
				dfs(s, i+1, k+1, path+"."+strconv.Itoa(t))
			}else{
				break
			}
		}
	}
}