/*
	给定一个可包含【重复数字】的序列，返回所有不重复的全排列
*/
/*
	方法1：回溯
	used数组去重
*/
var ans [][]int
var path []int
var nums []int
var used []bool

func permuteUnique(_nums []int) [][]int {
	// 😅😅😅 去重一定要对元素经行排序
	nums = BubbleSort(_nums)
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
		// used[i - 1] == true，说明【同一树支】nums[i - 1]使用过
		// used[i - 1] == false，说明【同一树层】nums[i - 1]使用过
		// 如果同一树层nums[i - 1]使用过则直接跳过
		// 😅😅😅 如果改成 used[i - 1] == true， 也是正确的。【对于排列问题，树层上去重和树枝上去重，都是可以的，但是树层上去重效率更高！】
		if i > 0 && nums[i] == nums[i-1] && used[i-1] == false {
			continue
		}
		// 😅 未被使用过
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

/*
	方法1：回溯
	😅😅😅 使用map去重
*/
var ans [][]int
var path []int
var nums []int
var used []bool

func permuteUnique(_nums []int) [][]int {
	nums = BubbleSort(_nums)
	used = make([]bool, len(nums))
	ans = make([][]int, 0)
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
	// 控制某一节点下的同一层元素不能重复
	usedMap := make(map[int]bool, 0)
	// 😅 每层都是从0开始搜索而不是start
	for i := 0; i < len(nums); i++ {
		if usedMap[nums[i]] {
			continue
		}
		if used[i] {
			continue
		}
		usedMap[nums[i]] = true
		used[i] = true
		path = append(path, nums[i])
		backtrack()
		// 回溯
		path = path[:len(path)-1]
		used[i] = false
	}
}