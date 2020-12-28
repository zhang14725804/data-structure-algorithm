var ans [][]int

func permute(nums []int) [][]int {
	path := make([]int, 0)
	backtrack(nums, path)
	return ans
}

func backtrack(nums []int, path []int) {
	// 结束条件
	if len(path) == len(nums) {
		ans = append(ans, path)
		return
	}
	for i := 0; i < len(nums); i++ {
		// 排除不合法的选择
		if contains(path, nums[i]) {
			continue
		}
		// 做选择
		path = append(path, nums[i])
		// 进入下一层决策树
		backtrack(nums, path)
		// 撤销选择
		path = path[:len(path)-1]
	}
}