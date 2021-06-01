/*
	给定一个字符串 s，将 s 分割成一些子串，使每个子串都是回文串。
	返回 s 所有可能的分割方案。

	思路：s[i,j]是回文串，那么s[i+1,j-1]也是回文串

*/

/*
	方法：回溯

	涉及到的问题（😅😅😅）：
	切割问题可以抽象为组合问题
	如何模拟那些切割线
	切割问题中递归如何终止
	在递归循环中如何截取子串
	如何判断回文

	其实切割问题类似组合问题
	递归用来【纵向遍历】，for循环用来【横向遍历】，切割线（就是图中的红线）切割到字符串的结尾位置，说明找到了一个切割方法。
*/
var ans [][]string
var path []string
var str string

func partition(s string) [][]string {
	str = s
	ans = make([][]string, 0) // 只是为了提交，leetcode提交时，ans 会拼接之前提交的结果
	backtrack(0)
	return ans
}

/*
	在处理组合问题的时候，递归参数需要传入 start ，表示下一轮递归遍历的起始位置，这个 start 就是切割线
*/
func backtrack(start int) {
	// 如果起始位置已经大于s的大小，说明已经找到了一组分割方案了
	if start >= len(str) {
		back := make([]string, len(path))
		copy(back, path)
		ans = append(ans, back)
		return
	}
	for i := start; i < len(str); i++ {
		// 是回文子串
		if isPalindrome(str[start : i+1]) {
			path = append(path, str[start:i+1])
		} else {
			// 不是回文，跳过
			continue
		}
		// 😅 切割过的位置，不能重复切割，所以，backtracking(s, i + 1); 传入下一层的起始位置为i + 1。
		backtrack(i + 1)
		path = path[:len(path)-1]
	}
}

func isPalindrome(str string) bool {
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		if str[i] != str[j] {
			return false
		}
	}
	return true
}