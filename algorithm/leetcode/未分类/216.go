/*
	找出所有相加之和为 n 的 k 个数的组合。组合中只允许含有 1 - 9 的正整数，并且每种组合中不存在重复的数字。

	todo:思路：回溯或者dfs
*/
var ans [][]int
var path []int
var k int

// todo:代码有问题，最后返回[[9,9,9]]，也就是说path这里有问题
func combinationSum3(_k int, n int) [][]int {
	k = _k
	helper(n, 1)
	return ans
}

func helper(n, start int) {
	if len(path) == k {
		if n == 0 {
			ans = append(ans, path)
		}
		return
	}

	for i := start; i < 10; i++ {
		path = append(path, i)
		helper(n-i, i+1)
		path = path[:len(path)-1]
	}
}