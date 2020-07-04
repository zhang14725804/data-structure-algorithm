/*
	给定 S 和 T 两个字符串，当它们分别被输入到空白的文本编辑器后，判断二者是否相等，并返回结果。 # 代表退格字符。
	注意：如果对空文本输入退格字符，文本继续为空。

	思路：两个指针，分别从S、T的尾部开始遍历数组，遇到退格字符，做标记，当前索引向前移动
	todo：测试可以通过，提交无法通过
*/
func backspaceCompare(S string, T string) bool {
	i,j := len(S)-1,len(T)-1
	// 退格键的数量
	bi,bj:=0,0
	for i>=0 || j>=0{
		for i>=0 {
			if S[i] == '#'{
				i--
				bi++
			}else if bi > 0{
				bi--
				i--
			}else{
				break
			}
		}
		for j>=0 {
			if T[j] == '#'{
				j--
				bj++
			}else if bj > 0{
				bj--
				j--
			}else{
				break
			}
		}
		if S[i] != T[j]{
			return false
		}
		if (i>=0 && j<0) || (j>=0 && i<0){
			return false
		}
		i--
		j--
	}
	return true
}


// 另一种思路
func backspaceCompare(S string, T string) bool {

    return Com(S) == Com(T)
}

func Com(s string) string {
    stack := make([]byte, 0)

    for i := 0; i < len(s); i++ {
		// 如果不是#，存入栈；如果是#，取出栈顶元素
        if s[i] != '#' {
            stack = append(stack, s[i])
        } else if len(stack) != 0{
            stack = stack[:len(stack)-1]
        }
    }
    return string(stack)
}
