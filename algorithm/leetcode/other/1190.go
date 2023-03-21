/*
	从左到右遍历该字符串，使用字符串 str 记录当前层所遍历到的字母,对于当前遍历的字符：
	1. 如果是左括号，将 str 插入到栈中，并将 str 置为空，进入下一层
	2. 如果是右括号，则说明遍历完了当前层，需要将 str 反转，返回给上一层。
	具体地，将栈顶字符串弹出，然后将反转后的 str 拼接到栈顶字符串末尾，将结果赋值给 str
	3. 如果是小写英文字母，将其加到 str 末尾

	对栈的理解不到位 😅😅😅
*/
func reverseParentheses(s string) string {
	stack := [][]byte{}
	cstr := []byte{}
	for i := range s {
		if s[i] == '(' {
			stack = append(stack, cstr)
			cstr = []byte{}
		} else if s[i] == ')' {
			// 反转当前括号中的内容
			for start, end := 0, len(cstr)-1; start < end; start, end = start+1, end-1 {
				cstr[end], cstr[start] = cstr[start], cstr[end]
			}
			// 😅 和当前栈顶序列拼接
			cstr = append(stack[len(stack)-1], cstr...)
			// 出栈
			stack = stack[:len(stack)-1]
		} else {
			cstr = append(cstr, s[i])
		}
	}
	return string(cstr)
}