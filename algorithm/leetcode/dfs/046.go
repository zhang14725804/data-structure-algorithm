/*
	46. Permutations：给定一个 没有重复 数字的序列，返回其所有可能的全排列。

	全排列问题，两种方法（这特么。。。）
	（1）枚举每个位置上放那个树
	（2）枚举每个数放在哪个位置

	回溯算法（精讲）：https://leetcode-cn.com/problems/permutations/solution/hui-su-suan-fa-python-dai-ma-java-dai-ma-by-liweiw/
*/

var (
	n    int     // 数组大小
	st   []bool  // 互斥锁
	ans  [][]int // 所有方法
	path []int   // 当前方案
)

func permute(nums []int) [][]int {
	// 方法一：枚举每个位置放什么数字
	n = len(nums)
	if n == 0 {
		return ans
	}
	st = make([]bool, n)
	dfs(nums, 0)
	return ans
}

func dfs(nums []int, u int) {
	// 已经便利完所有方案
	if u == n {
		// 这里是个什么坑（question）。直接ans = append(ans, path)，这种结果：[[3 2 1] [3 2 1] [3 2 1] [3 2 1] [3 2 1] [3 2 1]]
		c := make([]int, n)
		copy(c, path)
		ans = append(ans, c)
		return
	}
	for i := 0; i < n; i++ {
		// 当前位置可以放这个数
		if !st[i] {
			st[i] = true
			path = append(path, nums[i])
			// 递归下一层
			dfs(nums, u+1)
			st[i] = false
			path = path[:len(path)-1]
		}
	}
}