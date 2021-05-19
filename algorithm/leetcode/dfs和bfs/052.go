/*
	52. N-Queens II(八皇后问题，类似全排列问题。这个有难度)

	思路：每行每列只存在一个

	依次枚举每一行皇后的位置
	（1）每一列只能有一个皇后
	（2）每条斜线（正反两个方向）上只能有一个皇后

	todos::测试通过，提交不通过
*/
var (
	ans,n int
	// 行，正反对角线
	col,d,ud []bool
)
func totalNQueens(_n int) int {
	n = _n
	col = make([]bool,n)
	d,ud = make([]bool,2*n),make([]bool,2*n)
	dfs(0)
	return ans
}

func dfs(u int){
	if n==u{
		ans+=1
		return
	}
	for i:=0;i<n;i++{
		if !col[i] && !d[u+i] && !ud[u-i+n]{
			col[i],d[u+i],ud[u-i+n] = true,true,true
			dfs(u+1)
			col[i],d[u+i],ud[u-i+n] = false,false,false
		}
	}
}