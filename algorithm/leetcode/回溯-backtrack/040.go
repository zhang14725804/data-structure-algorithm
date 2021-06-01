import "fmt"

/*
	给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
	candidates 中的每个数字在每个组合中只能使用一次。

	说明：

	所有数字（包括目标数）都是正整数。
	解集不能包含重复的组合。
*/

/*
	方法：回溯+剪枝+ used数组去重

	都知道组合问题可以抽象为树形结构，那么“使用过”在这个树形结构上是有两个维度的，一个维度是同一树枝上“使用过”，一个维度是同一树层上“使用过”。
	回看一下题目，元素在同一个组合内是可以重复的，怎么重复都没事，但两个组合不能相同。
	所以我们要去重的是同一树层上的“使用过”，同一树枝上的都是一个组合里的元素，不用去重。

	question 😅😅😅 还是不懂去重套路
*/
var ans [][]int
var path []int
var candidates []int
var used []bool

func combinationSum2(_candidates []int, target int) [][]int {
	candidates = BubbleSort(_candidates) // 数组排序
	used = make([]bool, len(candidates)) // 数组中的数字是否使用过
	ans = make([][]int, 0)               // 只是为了提交，leetcode提交时，ans 会拼接之前提交的结果
	backtrack(target, 0)
	return ans
}

/*
	sum：目标和和已经收集到的元素和的差值。😅😅😅 这个参数有故事：targetSum和sum差值
	start：start来控制for循环的起始位置
	used：数组中的数字是否使用过
*/
func backtrack(sum int, start int) {
	// base case
	if sum == 0 {
		c := make([]int, len(path))
		copy(c, path)
		ans = append(ans, c)
	}
	// 递归+剪枝
	for i := start; i < len(candidates) && sum-candidates[i] >= 0; i++ {
		// 如果candidates[i] == candidates[i - 1] 并且 used[i - 1] == false，就说明：前一个树枝，使用了candidates[i - 1]，也就是说同一树层使用过candidates[i - 1]。
		// used[i - 1] == true，说明同一树支candidates[i - 1]使用过 😅😅😅
		// used[i - 1] == false，说明同一树层candidates[i - 1]使用过 😅😅😅
		if i > 0 && candidates[i] == candidates[i-1] && used[i-1] == false {
			continue
		}
		used[i] = true
		path = append(path, candidates[i])
		backtrack(sum-candidates[i], i+1)
		used[i] = false
		path = path[:len(path)-1]
	}
}

/*
	回溯+剪枝
	question 😅😅😅 直接用start来去重也是可以的， 就不用used数组
*/
var ans [][]int
var path []int
var candidates []int

func combinationSum2(_candidates []int, target int) [][]int {
	candidates = BubbleSort(_candidates)
	backtrack(target, 0)
	return ans
}

func backtrack(sum int, start int) {
	if sum == 0 {
		c := make([]int, len(path))
		copy(c, path)
		ans = append(ans, c)
	}
	for i := start; i < len(candidates) && sum-candidates[i] >= 0; i++ {
		// 😅 【i > start】 要对同一树层使用过的元素进行跳过
		if i > start && candidates[i] == candidates[i-1] {
			continue
		}
		path = append(path, candidates[i])
		// 这里是i+1，每个数字在每个组合中只能使用一次
		backtrack(sum-candidates[i], i+1)
		path = path[:len(path)-1]
	}
}

/*
	方法1：回溯
	😅😅😅 使用map去重
*/
var ans [][]int
var path []int
var candidates []int

func combinationSum2(_candidates []int, target int) [][]int {
	candidates = BubbleSort(_candidates) // 数组排序
	ans = make([][]int, 0)
	backtrack(target, 0)
	return ans
}

/*
	sum：目标和和已经收集到的元素和的差值。😅😅😅 这个参数有故事：targetSum和sum差值
	start：start来控制for循环的起始位置
*/
func backtrack(sum int, start int) {
	// base case
	if sum == 0 {
		c := make([]int, len(path))
		copy(c, path)
		ans = append(ans, c)
	}
	usedMap := make(map[int]bool, 0)
	// 递归+剪枝
	for i := start; i < len(candidates) && sum-candidates[i] >= 0; i++ {
		if usedMap[candidates[i]] {
			continue
		}
		usedMap[candidates[i]] = true
		path = append(path, candidates[i])
		backtrack(sum-candidates[i], i+1)
		path = path[:len(path)-1]
		// 😅😅😅 这里不需要回溯
	}
}