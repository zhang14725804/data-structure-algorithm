/*
	给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
	给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
*/

/*
	方法1：回溯

	从示例上来说，输入"23"，最直接的想法就是两层for循环遍历了吧，正好把组合的情况都输出了。
	如果输入"233"呢，那么就三层for循环，如果"2333"呢，就四层for循环.......
	😅😅😅 这for循环的层数如何写出来，此时又是回溯法登场的时候了。

	本题要解决如下三个问题：
	（1）数字和字母如何映射
	（2）两个字母就两个for循环，三个字符我就三个for循环，以此类推，然后发现代码根本写不出来
	（3）输入1 * #按键等等异常情况
*/
var ans []string
var digits string
var path string
var letterMap []string

func letterCombinations(_digits string) []string {
	// 判空
	if len(_digits) == 0 {
		return ans
	}
	digits = _digits
	letterMap = []string{
		"",     // 0
		"",     // 1
		"abc",  // 2
		"def",  // 3
		"ghi",  // 4
		"jkl",  // 5
		"mno",  // 6
		"pqrs", // 7
		"tuv",  // 8
		"wxyz", // 9
	}
	backtrack(0)
	return ans
}

/*
	index是记录遍历第几个数字了 😅😅😅
*/
func backtrack(index int) {
	if index == len(digits) {
		ans = append(ans, path)
		return
	}
	// 取出index位置的数字，所对应的字母
	c := digits[index]
	letters := letterMap[c-'0']
	for i := 0; i < len(letters); i++ {
		path += string(letters[i])
		backtrack(index + 1)
		// 回溯
		path = path[:len(path)-1]
	}
}

/*
	方法2：递归（隐藏着回溯）
*/
var res []string
var digits string
var nums []string

func letterCombinations(_digits string) []string {
	digits = _digits
	res = make([]string, 0)
	if digits == "" {
		return res
	}
	nums = []string{
		"",
		"",
		"abc",
		"def",
		"ghi",
		"jkl",
		"mno",
		"pqrs",
		"tuv",
		"wxyz",
	}

	dfs(0, "")
	return res
}

func dfs(index int, s string) {
	if index == len(digits) {
		res = append(res, s)
		return
	}
	c := digits[index]
	letters := nums[c-'0']
	for i := 0; i < len(letters); i++ {
		// 注意：这里递归隐藏着回溯 😅😅😅
		dfs(index+1, s+string(letters[i]))
	}
}