/*
	数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
	括号问题的一个性质：
		对于一个「合法」的括号字符串组合 p，必然对于任何 0 <= i < len(p) 都有：子串 p[0..i] 中左括号的数量都大于或等于右括号的数量。

	算法输入一个整数 n，让你计算 n 对儿括号能组成几种合法的括号组合，可以改写成如下问题：
	现在有 2n 个位置，每个位置可以放置字符 ( 或者 )，组成的所有括号组合中，有多少个是合法的？

	有关括号问题，你只要记住以下性质（🔥🔥🔥），思路就很容易想出来：
	1、一个「合法」括号组合的左括号数量一定等于右括号数量，这个很好理解。
	2、对于一个「合法」的括号字符串组合 p，必然对于任何 0 <= i < len(p) 都有：子串 p[0..i] 中左括号的数量都大于或等于右括号的数量。

*/
var ans []string
var path string

func generateParenthesis(n int) []string {
	if n == 0 {
		return ans
	}
	backtrack(n, n)
	return ans
}

// 用 left 记录还可以使用多少个左括号，用 right 记录还可以使用多少个右括号
func backtrack(left, right int) {
	// 非法方案; 左括号的数量都大于或等于右括号的数量
	if right < left || left < 0 || right < 0 {
		return
	}

	// 当所有括号都恰好用完时，得到一个合法的括号组合
	if left == 0 && right == 0 {
		ans = append(ans, path)
		return
	}

	// 选择，尝试放一个左括号
	path += "("
	// 回溯
	backtrack(left-1, right)
	// 撤销选择
	path = path[:len(path)-1]

	// 选择，尝试放一个右括号
	path += ")"
	// 回溯
	backtrack(left, right-1)
	// 撤销选择
	path = path[:len(path)-1]
}