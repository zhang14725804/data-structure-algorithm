/*
	删除最小数量的无效括号，使得输入的字符串有效，返回所有可能的结果。
	
	比较难

	合法括号序列两个性质：
		（1）左右括号数量相同
		（2）任意前缀中左括号数量>=右括号数量

*/
var ans []string
func removeInvalidParentheses(s string) []string {
	l:=0 // 当前左括号数量减去右括号数量，要删除的左括号数量
	r:=0 // 当前需要删除多少个右括号

	for _,c:=range s{
		if c=='(' {
			l++
		}else if c==')'{
			if l==0{ // 做右括号数量相同,需要删除右括号
				r++
			}else{
				l-- 
			}
		}
	}
	
	dfs(s,0,"",0,l,r)
	return ans
}
/*
	todo：这个dfs很难理解
	
	字符串s
	当前枚举到那个字符
	当前字符串是什么（当前删除之后剩余的字符串）
	当前左括号减去右括号的数量
	当前可以删除多少个左括号
	当前可以删除多少个右括号
*/
func dfs(s string,u int,path string,cnt int,l int,r int){
	// 已经递归到字符串结尾
	if u == len(s){
		if cnt == 0{
			ans = append(ans,path)
		}
		return
	}
	
	if s[u]!='(' && s[u]!=')'{ // 不是左右括号的情况
		// 搜索下一个字符
		dfs(s,u+1,path+string(s[u]),cnt,l,r)
	}else if s[u]=='('{ 
		k:=u
		// 连续左括号的情况
		for k<len(s) && s[k]=='('{  
			k++
		}
		l-=k-u // k-u当前连续左括号的数量
		// 
		for i:=k-u;i>=0;i--{
			// 删掉的左括号数量没有超过限制
			if l>=0{
				dfs(s,k,path,cnt,l,r)
			}
			path += string('(')
			cnt++
			l++
		}
	}else if s[u]==')'{
		k:=u
		for k<len(s) && s[k]==')'{
			k++
		}
		r-=k-u
		for i:=k-u;i>=0;i--{
			if cnt>=0 && r>=0{
				dfs(s,k,path,cnt,l,r)
			}
			path+=string(')')
			cnt--
			r++
		}
	}
}