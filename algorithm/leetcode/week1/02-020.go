/*
	给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效
	有效字符串需满足：
		左括号必须用相同类型的右括号闭合。
		左括号必须以正确的顺序闭合。
		注意空字符串可被认为是有效字符串。

	思路：遍历字符串，存入栈中，和栈顶元素进行比较。思路经典
*/
func isValid(s string) bool {
	stk := &stack{}
	// 注意：这里c是rune类型
	for _,c:=range s{
		if c == ')'{
			if stk.empty() || stk.top() != '('{
				return false
			}
			stk.pop()
		}else if c == ']'{
			if stk.empty() || stk.top() != '['{
				return false
			}
			stk.pop()
		}else if c == '}'{
			if stk.empty() || stk.top() != '{'{
				return false
			}
			stk.pop()
		}else{
			stk.push(c)
		}
	}
	return stk.empty()
}

// 栈
type stack struct{
	x []rune
}
// 入栈
func (this *stack) push(x rune){
	this.x = append(this.x,x)
}
// 出栈
func (this *stack) pop(){
	this.x= this.x[:len(this.x)-1]
}
// 返回栈顶元素
func (this *stack) top() rune{
	return this.x[len(this.x)-1]
}
// 是否为空
func (this *stack) empty() bool{
	return len(this.x) == 0
}