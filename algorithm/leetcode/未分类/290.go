/*
	给定一种规律 pattern 和一个字符串 str ，判断 str 是否遵循相同的规律。
	这里的 遵循 指完全匹配，例如， pattern 里的每个字母和字符串 str 中的每个非空单词之间存在着双向连接的对应规律。

	不会😅（todo）
*/
/*
	方法：双射（即使单射又是满射）
	离散数学集合论？（todo）
	满射：位置一一对应
	单设：a1和a2各自对应b1,b2
*/
func wordPattern(pattern string, s string) bool {
	words = strings.Split(s, " ")
	if len(words) != len(pattern) {
		return false
	}

	// pattern<=>words 的映射
	pw := make(map[byte]string, 0)
	wp := make(map[string]byte, 0)
	for i := 0; i < len(pattern); i++ {
		a := pattern[i]
		b := words[i]

		if v, ok := pw[a]; ok && v != b {
			return false
		}
		pw[a] = b
		if v, ok := wp[b]; ok && v != a {
			return false
		}
		wp[b] = a
	}
	return true
}