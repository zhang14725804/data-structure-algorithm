/*
	方法1：回溯
	😅😅😅 深度优先遍历、递归、栈，它们三者背后统一的逻辑都是「后进先出」
	「回溯算法」强调了「深度优先遍历」思想的用途，用一个 不断变化 的变量，在尝试各种可能的过程中，搜索需要的结果。强调了 回退 操作对于搜索的合理性
	回溯算法用于 搜索一个问题的所有的解 ，通过深度优先遍历的思想实现。
*/
var ans [][]int // 结果
var path []int  // 当前路径
var nums []int  // 参数
var nLen int    // 参数长度
var used []bool // 😅 标记已经用过的元素

func permute(_nums []int) [][]int {
	nums = _nums
	nLen = len(nums)
	used = make([]bool, nLen)
	ans = make([][]int, 0)
	backtrack()
	return ans
}

func backtrack() {
	// （1） 😅 base case，递归出口
	if len(path) == nLen {
		back := make([]int, len(path))
		copy(back, path)
		ans = append(ans, back)
		// 记得返回
		return
	}
	// （2） 每层都是从0开始搜索而不是start
	for i := 0; i < nLen; i++ {
		if used[i] {
			continue
		}

		// （3）递归
		used[i] = true
		path = append(path, nums[i])

		backtrack()

		// （4）逆向操作
		path = path[:len(path)-1]
		used[i] = false
	}
}