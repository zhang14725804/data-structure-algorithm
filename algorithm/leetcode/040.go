import "fmt"

/*
	给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

	candidates 中的每个数字在每个组合中只能使用一次。

	说明：

	所有数字（包括目标数）都是正整数。
	解集不能包含重复的组合。

	todo:看着好像没有问题，但是测试不通过
*/
var ans = make([][]int, 0)
var path = make([]int, 0)

func combinationSum2(c []int, target int) [][]int {
	c = quickSort(c)
	fmt.Println(c)
	dfs(c, 0, target)
	return ans
}
func dfs(c []int, u int, target int) {
	if target == 0 {
		// 坑在这里
		c := make([]int, len(path))
		copy(c, path)
		ans = append(ans, c)
		return
	}
	if u == len(c) {
		return
	}
	// 相同的元素
	k := u + 1
	for k < len(c) && c[k] == c[u] {
		k++
	}
	cnt := k - u

	for i := 0; i <= cnt && c[u]*i <= target; i++ {
		dfs(c, k, target-c[u]*i)
		path = append(path, c[u])
	}
	// 恢复现场
	for i := 0; i <= cnt && c[u]*i <= target; i++ {
		path = path[:len(path)-1]
	}
}
