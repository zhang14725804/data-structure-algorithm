/*
	给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。
*/

/*
	方法1：回溯
	😅😅😅
	直接的解法当然是使用for循环，例如示例中k为2，很容易想到 用两个for循环，这样就可以输出 和示例中一样的结果。
	输入：n = 100, k = 3 那么就三层for循环
	如果n为100，k为50呢，那就50层for循环，是不是开始窒息。
	此时就会发现虽然想暴力搜索，但是用for循环嵌套连暴力都写不出来！

	要解决 n为100，k为50的情况，暴力写法需要嵌套50层for循环，那么回溯法就用递归来解决嵌套层数的问题。
	递归来做层叠嵌套（可以理解是开k层for循环），每一次的递归中嵌套一个for循环，那么递归就可以用于解决多层嵌套循环的问题了。

	回溯法解决的问题都可以抽象为树形结构（N叉树），用树形结构来理解回溯就容易多了
	每次从集合中选取元素，可选择的范围随着选择的进行而收缩，调整可选择的范围。
	n相当于树的宽度，k相当于树的深度。
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

/*
	start：😅 每次从集合中选取元素，可选择的范围随着选择的进行而收缩，调整可选择的范围，就是要靠 start
*/
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
	😅😅😅 【剪枝】 😅😅😅
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