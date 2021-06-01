/*
	找出所有相加之和为 n 的 k 个数的组合。组合中只允许含有 1 - 9 的正整数，并且每种组合中不存在重复的数字。
*/

/*
	方法1：回溯
	【k】相当于树的深度
	【1~9】相当于树的宽度
*/
var ans [][]int
var path []int
var k int

func combinationSum3(_k int, n int) [][]int {
	k = _k
	ans = make([][]int, 0) // 只是为了提交，leetcode提交时，ans 会拼接之前提交的结果
	backtrack(n, 1)
	return ans
}

/*
	start，下一层for循环搜索的起始位置
	n：目标和和已经收集到的元素和的差值。😅😅😅 这个参数有故事：targetSum和sum差值
*/
func backtrack(n, start int) {
	if n == 0 && len(path) == k {
		c := make([]int, k)
		copy(c, path)
		ans = append(ans, c)
		return
	}
	for i := start; i < 10; i++ {
		path = append(path, i)
		backtrack(n-i, i+1)
		path = path[:len(path)-1]
	}
}

/*
	😅😅😅 【剪枝】
*/
func backtrack(n, start int) {
	// 剪枝操作
	if n < 0 {
		return
	}
	if n == 0 && len(path) == k {
		c := make([]int, k)
		copy(c, path)
		ans = append(ans, c)
		return
	}
	// for循环的范围也可以剪枝，【i < 10 - (k - path.size()) + 1】 就可以了
	for i := start; i < 10-(k-len(path))+1; i++ {
		path = append(path, i)
		backtrack(n-i, i+1)
		path = path[:len(path)-1]
	}
}