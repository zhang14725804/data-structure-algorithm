/*
	给定一个**没有重复**数字的序列，返回其所有可能的全排列。
*/

/*
	方法1：回溯
*/
var ans [][]int
var path []int
var nums []int
var used []bool // 😅 需要used数组记录path里都放了哪些元素了

func permute(_nums []int) [][]int {
	nums = _nums
	used = make([]bool, len(nums))
	backtrack()
	return ans
}

func backtrack() {
	// 😅 base case，递归出口
	if len(path) == len(nums) {
		back := make([]int, len(path))
		copy(back, path)
		ans = append(ans, back)
		return
	}
	// 😅 每层都是从0开始搜索而不是start
	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}
		used[i] = true
		path = append(path, nums[i])
		backtrack()
		// 回溯
		path = path[:len(path)-1]
		used[i] = false
	}
}