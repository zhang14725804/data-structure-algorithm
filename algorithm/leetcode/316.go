/*
	给你一个字符串 s ，请你去除字符串中重复的字母，使得每个字母只出现一次。需保证 返回结果的字典序最小（要求不能打乱其他字符的相对位置）。

	要求一、要去重。
	要求二、去重字符串中的字符顺序不能打乱 s 中字符出现的相对顺序。
	要求三、在所有符合上一条要求的去重字符串中，字典序最小的作为最终结果。（难点在这里）

	316和1081是同一題
*/
func removeDuplicateLetters(s string) string {
	stk := make([]byte, 0)    // 单调栈的思路
	count := make([]int, 256) // 统计每个单词出现的次数，为之后是否可以删除栈顶元素做铺垫
	for i := 0; i < len(s); i++ {
		count[s[i]]++
	}

	inStack := make([]bool, 256) // 是否在栈中，位圖法
	for i := 0; i < len(s); i++ {
		c := s[i]
		count[c]-- // 每次遍历一次，对应的计数减1
		// 如果字符 c 存在栈中，直接跳过
		if inStack[c] {
			continue
		}
		// 插入之前，和栈顶元素比较一下大小；如果字典序比前面的小，pop 前面的元素
		for len(stk) > 0 && stk[len(stk)-1] > c {
			// 再不存在 栈顶元素，停止pop
			if count[stk[len(stk)-1]] == 0 {
				break
			}
			// 若之后还有，则可以 pop
			inStack[stk[len(stk)-1]] = false
			stk = stk[:len(stk)-1]
		}
		// 若不存在，则插入栈顶并标记为存在
		stk = append(stk, c)
		inStack[c] = true
	}
	return string(stk)
}