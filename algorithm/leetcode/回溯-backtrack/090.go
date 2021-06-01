/*
	给你一个整数数组 nums ，其中可能包含【重复元素】，请你返回该数组所有可能的子集（幂集）。
	解集 不能 包含重复的子集。返回的解集中，子集可以按 任意顺序 排列。
*/

/*
	方法1：回溯
	😅😅😅 使用used数组去重
*/
var ans [][]int
var path []int
var used []bool
var nums []int

func subsetsWithDup(_nums []int) [][]int {
	// 😅 去重一定要对元素经行排序
	nums = BubbleSort(_nums)
	used = make([]bool, len(nums))
	backtrack(0)
	return ans
}

func backtrack(start int) {
	back := make([]int, len(path))
	copy(back, path)
	ans = append(ans, back)

	for i := start; i < len(nums); i++ {
		// used[i - 1] == true，说明同一树支 nums[i - 1]使用过
		// used[i - 1] == false，说明同一树层 nums[i - 1]使用过
		// 而我们要对同一树层使用过的元素进行跳过
		if i > 0 && used[i-1] == false && nums[i] == nums[i-1] {
			continue
		}
		used[i] = true
		path = append(path, nums[i])
		backtrack(i + 1)
		used[i] = false
		path = path[:len(path)-1]
	}
}

/*
	本题也可以不用used数组来去重，因为递归的时候下一个start是i+1而不是0
	question 😅😅😅 直接用start来去重也是可以的， 就不用used数组
*/

var ans [][]int
var path []int
var nums []int

func subsetsWithDup(_nums []int) [][]int {
	nums = BubbleSort(_nums)
	ans = make([][]int, 0) // 只是为了提交，leetcode提交时，ans 会拼接之前提交的结果
	backtrack(0)
	return ans
}

func backtrack(start int) {
	back := make([]int, len(path))
	copy(back, path)
	ans = append(ans, back)

	for i := start; i < len(nums); i++ {
		// 😅😅😅 【i > start】 要对同一树层使用过的元素进行跳过
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		path = append(path, nums[i])
		backtrack(i + 1)
		path = path[:len(path)-1]
	}
}

/*
	方法1：回溯
	😅😅😅 使用map去重
*/
var ans [][]int
var path []int
var nums []int

func subsetsWithDup(_nums []int) [][]int {
	nums = BubbleSort(_nums)
	backtrack(0)
	return ans
}

func backtrack(start int) {
	back := make([]int, len(path))
	copy(back, path)
	ans = append(ans, back)
	// 定义usedMap对同一节点下的【本层去重】
	usedMap := make(map[int]bool, 0)
	for i := start; i < len(nums); i++ {
		// 如果发现出现过就pass
		if usedMap[nums[i]] {
			continue
		}
		// usedMap更新元素
		usedMap[nums[i]] = true
		path = append(path, nums[i])
		backtrack(i + 1)
		path = path[:len(path)-1]
		// 😅😅😅 这里不需要回溯
	}
}