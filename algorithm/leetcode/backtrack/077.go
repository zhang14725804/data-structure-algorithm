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

func backtrack(start int) {
	// 如果 temp 里的数字够了 k 个，就把它加入到结果中
	if len(path) == k {
		temp := make([]int, k)
		copy(temp, path)
		ans = append(ans, temp)
	}
	for i := start; i <= n; i++ {
		path = append(path, i)
		backtrack(i + 1) // i+1 而不是start+1
		path = path[:len(path)-1]
	}
}