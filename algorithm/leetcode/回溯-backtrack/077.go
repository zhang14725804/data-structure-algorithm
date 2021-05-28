/*
	给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。
*/
/*
	方法1：回溯
*/
var ans [][]int
var path []int
var n, k int

func combine(_n int, _k int) [][]int {
	n, k = _n, _k
	// 😅 为什么要有这个start呢？
	backtrack(1)
	return ans
}

// 😅 每次从集合中选取元素，可选择的范围随着选择的进行而收缩，调整可选择的范围，就是要靠 start
func backtrack(start int) {
	// base case，终止条件
	if len(path) == k {
		temp := make([]int, k)
		copy(temp, path)
		ans = append(ans, temp)
		return
	}
	for i := start; i <= n; i++ {
		// 选择
		path = append(path, i)
		// 进入下一层决策树；i+1 而不是start+1
		backtrack(i + 1)
		// 取消选择
		path = path[:len(path)-1]
	}
}

/*
	方法1：回溯+剪枝
	重点在剪枝
*/
func backtrack(start int) {
	if len(path) == k {
		temp := make([]int, k)
		copy(temp, path)
		ans = append(ans, temp)
		return
	}
	/*
		可以剪枝的地方就在递归中每一层的for循环所选择的起始位置。
		如果for循环选择的起始位置之后的元素个数 已经不足 我们需要的元素个数了，那么就没有必要搜索了。
		来举一个例子，n = 4，k = 4的话，那么第一层for循环的时候，从元素2开始的遍历都没有意义了。 在第二层for循环，从元素3开始的遍历都没有意义了。

		（1）已经选择的元素个数：path.size();
		（2）还需要的元素个数为: k - path.size();
		（3）在集合n中至多要从该起始位置 : n - (k - path.size()) + 1，开始遍历
	*/
	for i := start; i <= n-(k-len(path))+1; i++ {
		path = append(path, i)
		backtrack(i + 1)
		path = path[:len(path)-1]
	}
}