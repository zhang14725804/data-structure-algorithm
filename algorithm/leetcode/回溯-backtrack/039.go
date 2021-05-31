/*
	给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
	candidates 中的数字可以无限制重复被选取。

	说明：
	所有数字（包括 target）都是正整数。
	解集不能包含重复的组合。
*/

/*
	方法1：回溯
	本题没有组合数量要求，仅仅是总和的限制，所以递归没有层数的限制，只要选取的元素总和超过target
*/
var candidates []int
var target int
var path []int
var ans [][]int

func combinationSum(_candidates []int, target int) [][]int {
	candidates = _candidates
	backtrack(target, 0)
	return ans
}

/*
	n：目标和和已经收集到的元素和的差值。😅😅😅 这个参数有故事：targetSum和sum差值
	start：start来控制for循环的起始位置
*/
func backtrack(n int, start int) {
	// base case
	if n < 0 {
		return
	}
	if n == 0 {
		c := make([]int, len(path))
		copy(c, path)
		ans = append(ans, c)
	}
	for i := start; i < len(candidates); i++ {
		path = append(path, candidates[i])
		// 😅😅😅 关键点:不用i+1了，表示可以重复读取当前的数
		backtrack(n-candidates[i], i)
		path = path[:len(path)-1]

		/*
			🐷🐷🐷 也可以这么写，这个backtrack中有两个回溯条件
			sum += candidates[i];
			path.push_back(candidates[i]);
			backtracking(candidates, target, sum, i); // 关键点:不用i+1了，表示可以重复读取当前的数
			sum -= candidates[i];   // 回溯
			path.pop_back();        // 回溯
		*/
	}
}

/*
	😅😅😅 【剪枝】

	对于sum已经大于target的情况，其实是依然进入了下一层递归，只是下一层递归结束判断的时候，会判断sum > target的话就返回。
	其实如果已经知道下一层的sum会大于target，就没有必要进入下一层递归了。
*/
func backtrack(n int, start int) {
	// base case
	if n < 0 {
		return
	}
	if n == 0 {
		c := make([]int, len(path))
		copy(c, path)
		ans = append(ans, c)
	}
	// 如果 n-candidates[i] < 0  就终止遍历
	for i := start; i < len(candidates) && n-candidates[i] >= 0; i++ {
		path = append(path, candidates[i])
		// 😅😅😅 关键点:不用i+1了，表示可以重复读取当前的数
		backtrack(n-candidates[i], i)
		path = path[:len(path)-1]
	}
}