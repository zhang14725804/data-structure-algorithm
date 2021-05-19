/*
	给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。
*/
var ans [][]int
var path []int
var n, k int

func combine(_n int, _k int) [][]int {
	n, k = _n, _k
	backtrack(1)
	return ans
}

// start 从第i个数字开始
func backtrack(start int) {
	// 结束条件 temp 里的数字够了 k 个
	if len(path) == k {
		temp := make([]int, k)
		copy(temp, path)
		ans = append(ans, temp)
		return
	}
	for i := start; i <= n; i++ {
		// 选择
		path = append(path, i)
		// 进入下一层决策树；i+1 而不是start+1
		backtrack(i + 1)
		// 取消选择
		path = path[:len(path)-1]
	}
}