
/*
	给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。

	这类问题如何表示坐标，是个问题
*/
func generate(n int) [][]int {
	f:=make([][]int,0)
	for i := 0; i < n; i++ {
		// 
		line:=make([]int,i+1)
		// 每行第一个和最后一个
		line[0]=1
		line[i]=1
		// 每列
		for j := 1; j < i; j++ {
			line[j]= f[i-1][j-1]+f[i-1][j]
		}
		f=append(f,line)
	}
	return f
}

