	/*
	给定一个**没有重复**数字的序列，返回其所有可能的全排列。
	0103 懵逼状态
*/

/*
	方法1：回溯
*/
var ans [][]int
var path []int
var nums []int
var nLen int
var used []bool // 😅 需要used数组记录path里都放了哪些元素了

func permute(_nums []int) [][]int {
	nums = _nums
	nLen = len(nums)
	used = make([]bool, nLen)
	ans = make([][]int, 0) 
	backtrack()
	return ans
}

func backtrack() {
	// （1） base case，递归出口
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
		// （3）
		used[i] = true
		path = append(path, nums[i])
		backtrack()
		path = path[:len(path)-1]
		used[i] = false
	}
}