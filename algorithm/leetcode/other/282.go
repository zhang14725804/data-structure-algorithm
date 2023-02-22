/*
	给表达式添加运算符

	方法：回溯法（todo：困难呀）
*/
var num string
var target int
var result []string

func addOperators(_num string, _target int) []string {
	num = _num
	target = _target
	backTrack("", 0, 0, 0)
	return result
}

/*
	path：当前路径
	start：起始位置
	eval：计算结果
	prev：记录上一个操作数
*/
func backTrack(path string, start int, eval int, prev int) {
	// 回溯出口
	if start == len(num) {
		if target == eval {
			result = append(result, path)
		}
		return
	}
	for i := start; i < len(num); i++ {
		// 数字不能以0开头
		if num[start] == '0' && i > start {
			continue
		}

		cur, _ := strconv.Atoi(num[start : i+1])

		if start == 0 {
			// 第一个数字不加符号
			backTrack(path+fmt.Sprint(cur), i+1, cur, cur)
		} else {
			//
			backTrack(path+"+"+fmt.Sprint(cur), i+1, eval+cur, cur)
			backTrack(path+"-"+fmt.Sprint(cur), i+1, eval-cur, -cur)
			// 先减去之前的值，然后加上 当前值和前边的操作数相乘（question）
			backTrack(path+"*"+fmt.Sprint(cur), i+1, eval-prev+prev*cur, prev*cur)
		}
	}
}