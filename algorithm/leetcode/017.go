/*
	question，我居然想到的只是循环😅
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

// 递归 digits位置和s字符串
func dfs(index int, s string) {
	if index == len(digits) {
		res = append(res, s)
		return
	}
	c := digits[index]
	letters := nums[c-'0']
	for i := 0; i < len(letters); i++ {
		dfs(index+1, s+string(letters[i]))
	}
}