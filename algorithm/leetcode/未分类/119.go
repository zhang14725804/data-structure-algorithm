/*
	给定一个非负索引 k，其中 k ≤ 33，返回杨辉三角的第 k 行。
	你可以优化你的算法到 O(k) 空间复杂度吗？

	思路1：普通二维数组
	思路2：滚动数组
*/

func getRow(n int) []int {
	f:=make([][]int,n+1)
	// i <= n
	for i := 0; i <= n; i++ {
		f[i] = make([]int,n+1)
	}
	// i <= n
	for i := 0; i <= n; i++ {
		f[i][0]=1
		f[i][i]=1
		// 
		for j := 1; j < i; j++ {
			f[i][j] = f[i-1][j-1]+f[i-1][j]
		}
	}
	return f[n]
}

// 思路2：滚动数组
func getRow(n int) []int {
	f:=make([][]int,n+1)
	// i <= n
	for i := 0; i <= n; i++ {
		f[i] = make([]int,n+1)
	}
	// i <= n
	for i := 0; i <= n; i++ {
		// todo:骚操作在这里；每行都&1
		f[i&1][0]=1
		f[i&1][i]=1
		// 
		for j := 1; j < i; j++ {
			// 注意优先级
			f[i&1][j] = f[(i-1)&1][j-1]+f[(i-1)&1][j]
		}
	}
	return f[n&1]
}