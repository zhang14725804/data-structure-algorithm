package leetcode

import "fmt"

/*
	所有 DNA 都由一系列缩写为 A，C，G 和 T 的核苷酸组成，例如：“ACGAATTCCG”。
	在研究 DNA 时，识别 DNA 中的重复序列有时会对研究非常有帮助。

	编写一个函数来查找目标子串，目标子串的长度为 10，且在 DNA 字符串 s 中出现次数超过一次。

	思路，循环字符串，每次前进10步，截取10个长度的字符串，存入hash表作为key，出现的次数作为value
*/
func Leetcode187() {
	s := "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
	// hash
	mp := make(map[string]int)
	// array
	var res []string
	// 注意临界条件
	for i := 0; i+10 <= len(s); i++ {
		key := s[i : i+10]
		mp[key]++
		if mp[key] == 2 {
			res = append(res, key)
		}
	}
	fmt.Println(res)
}
