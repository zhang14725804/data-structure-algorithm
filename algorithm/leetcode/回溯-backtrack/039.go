/*
	给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
	candidates 中的数字可以无限制重复被选取。

	说明：
	所有数字（包括 target）都是正整数。
	解集不能包含重复的组合。

	思路：暴力搜索（考虑搜索顺序）

	todo:看着好像没有问题，但是测试不通过
*/
var ans = make([][]int, 0)
var path = make([]int, 0)

func combinationSum(c []int, target int) [][]int {
	dfs(c, 0, target)
	return ans
}

// u 当前枚举到的位置
func dfs(c []int, u int, target int) {
	// 满足要求
	if target == 0 {
		// 这里path是对的
		ans = append(ans, path)
		return
	}
	// 已经枚举到最后一个数
	if u == len(c) {
		return
	}
	// 枚举取每个数字的个数
	for i := 0; c[u]*i <= target; i++ {
		// 枚举下一个位置
		dfs(c, u+1, target-c[u]*i)
		// 每取一个数字，都添加到当前集合
		path = append(path, c[u])
	}
	// 恢复现场
	for i := 0; c[u]*i <= target; i++ {
		path = path[:len(path)-1]
	}
}