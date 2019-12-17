package leetcode

import (
	"fmt"
	"math"
	"strings"
)

/*
	给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

	示例 1:

	输入: "abcabcbb"
	输出: 3
	解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
	示例 2:

	输入: "bbbbb"
	输出: 1
	解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
	示例 3:

	输入: "pwwkew"
	输出: 3
	解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
	     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

	先假设start不动，那么在理想的情况下，我们希望可以一直右移j，直到j到达原字符串的结尾，此时j-start就构成了一个候选的最长子串。
	每次都维护一个maxLen，就可以选出最长子串了。

	实际情况是，不一定可以一直右移j，如果字符j已经重复出现过（假设在位置k），就需要停止右移了。
	记录当前的候选子串并和maxLen做比较。接下来为下一次搜寻做准备。

	在下一次搜寻中，start应该更新到k+1。这句话的意思是，用这个例子来理解，
	abcdef是个不重复的子串，abcdefc中（为了方便记录为abc1defc2）,c1和c2重复了。
	那么下一次搜寻，应该跨过出现重复的地方进行，否则找出来的候选串依然有重复字符，
	且长度还不如上次的搜索。所以下一次搜索，直接从c1的下一个字符d开始进行，也就是说，下一次搜寻中，start应该更新到k+1
*/
func Letcode003(s string) int {
	if s == "" {
		return 0
	}
	// 字符串转数组
	slice := strings.Split(s, "")
	sliceLen := len(slice)
	// 为什么要float64（math.Max参数和返回值都是float64）
	maxLen := float64(0)
	// 已经扫描过的字符
	exit := make(map[string]int)
	// 使用i和j两个指针进行搜索，i代表候选的最长子串的开头，j代表候选的最长子串的结尾。
	end, start := 0, 0
	for {
		if end == sliceLen {
			break
		}
		if _, ok := exit[slice[end]]; ok && exit[slice[end]] >= start {
			maxLen = math.Max(maxLen, float64(end-start))
			// i应该更新到k+1
			start = exit[slice[end]] + 1
			exit[slice[start]] = start
			if end == start {
				end++
			}
		} else {
			exit[slice[end]] = end
			end++
		}
	}
	maxLen = math.Max(float64(maxLen), float64(end-start))
	fmt.Println(exit)
	fmt.Println(int(maxLen))
	return int(maxLen)
}
