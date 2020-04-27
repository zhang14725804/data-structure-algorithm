/*
	正则表达式匹配
	又是动态规划😅
	todo："aa"，"b*"测试未通过
	字符串"aaa"与模式"a.a"和"ab*ac*a"匹配，但是与"aa.a"和"ab*a"均不匹配 测试通过
	""和".*"测试通过（acwing提供的编译器好像有问题）
*/
var n,m int
var s,p string
var f [][]bool
func isMatch(_s string, _p string) bool {
	s,p = _s,_p
	n,m = len(s),len(p)
	f:=make([][]bool,n+1)
	for i:=0;i<n+1;i++{
		f[i] = make([]bool,m+1)
	}
	return dp(0,0,f)
}

func dp(x,y int,f [][]bool) bool{

	if f[x][y] != false{
		return f[x][y]
	}
	if y == m{
		b:=false
		if x == n{
			b = true
		}
		f[x][y] = b
		return f[x][y]
	}

	firstMatch := x < n && (p[y] == '.' || s[x] == p[y])
	if y+1 < m && p[y+1] == '*'{
		b:=false
		if dp(x, y+2,f) || dp(x+1, y,f) {
			b = true
		}
		f[x][y] = b
	}else{
		f[x][y] = firstMatch && dp(x+1, y+1,f)
	}
	return f[x][y]
}