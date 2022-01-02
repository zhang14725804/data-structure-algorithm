/*
	给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效
	有效字符串需满足：
		左括号必须用相同类型的右括号闭合。
		左括号必须以正确的顺序闭合。
		注意空字符串可被认为是有效字符串。

	思路：遍历字符串，存入【栈】中，和栈顶元素进行比较。
	0102 懵逼状态
*/
func isValid(s string) bool {
    var stack []byte
    for i:=0;i<len(s);i++{
		c:=s[i]
		// （1）左括号直接入栈
        if c == '(' ||  c == '['||  c == '{'{
            stack = append(stack,c)
        }else{
			// （2）右括号判断，是否和栈顶括号匹配。
			// 若匹配出栈继续判断下一个；不匹配直接返回false
            if len(stack)>0 && leftOff(c) == stack[len(stack)-1]{
				// 出栈
                stack = stack[:len(stack)-1]
            }else{
                return false
            }
        }
    }
    return len(stack) == 0
}

func leftOff(c byte) byte{
    if c == '}'{
        return '{'
    }
    if c == ']'{
        return '['
    }
    return '('
}