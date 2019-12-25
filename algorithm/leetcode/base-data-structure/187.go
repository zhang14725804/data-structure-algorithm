package leetcode

import "fmt"

// Leetcode187 Repeated DNA Sequences
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
