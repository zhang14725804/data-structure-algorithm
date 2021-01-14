/*
	请你来实现一个 atoi 函数，使其能将字符串转换成整数。

	三种方法：
		（1）硬编码
		（2）状态机
		（3）正则表达式

*/

/*
	方法1：有限状态机（question😅），机智啊

	字符串处理的题目往往涉及复杂的流程以及条件情况，如果直接上手写程序，一不小心就会写出极其臃肿的代码。

	因此，为了有条理地分析每个输入字符的处理方法，我们可以使用自动机这个概念：

	我们的程序在每个时刻有一个状态 s，每次从序列中输入一个字符 c，并根据字符 c 转移到下一个状态 s'。这样，我们只需要建立一个覆盖所有情况的从 s 与 c 映射到 s' 的表格即可解决题目中的问题。

*/
func myAtoi(str string) int {
	auto := &Automaton{
		Sign:  1,
		Ans:   0,
		state: "start",
		table: map[string][]string{
			"start":     []string{"start", "signed", "in_number", "end"},
			"signed":    []string{"end", "end", "in_number", "end"},
			"in_number": []string{"end", "end", "in_number", "end"},
			"end":       []string{"end", "end", "end", "end"},
		},
	}
	for i := 0; i < len(str); i++ {
		auto.get(str[i])
	}
	return auto.Sign * auto.Ans
}

var MAX int = (1 << 31) - 1
var MIN int = -(1 << 31)

type Automaton struct {
	Sign  int
	Ans   int
	state string
	table map[string][]string
}

// 这里需要传递指针（要改变struct的值）
func (auto *Automaton) get(c byte) {
	auto.state = auto.table[auto.state][auto.getCol(c)]
	if "in_number" == auto.state {
		auto.Ans = auto.Ans*10 + int(c-'0')
		// 极值处理
		if auto.Sign == 1 {
			auto.Ans = MinInt(MAX, auto.Ans)
		} else {
			auto.Ans = MinInt(-MIN, auto.Ans)
		}
	} else if "signed" == auto.state {
		if c == '+' {
			auto.Sign = 1
		} else {
			auto.Sign = -1
		}
	}
}

// 获取状态
func (auto Automaton) getCol(c byte) int {
	if c == ' ' {
		return 0
	}
	if c == '+' || c == '-' {
		return 1
	}
	if c >= '0' && c <= '9' {
		return 2
	}
	return 3
}

/*
	方法2：正则表达式（todo）
*/
func myAtoi(str string) int {
	arr := regexp.MustCompile(`[\+\-]?\d+`).FindAllStringSubmatch(str, -1)
	if arr[0] != nil {
		res := 0
		flag := 1
		if arr[0][0] == "0" {
			return res
		}
		if arr[0][0] == "-" {
			flag = -1
		}
		for _, c := range arr[0] {
			if res > 0 {
				res *= 10
			}
			// cannot use '0' (type untyped rune) as type string (solution.go)
			res += int(c - '0')
		}
		if flag == -1 {
			return -res
		}
		return res
	}
	return 0
}

// 自己的方法，是错的😅
func myAtoi1(str string) int {
	// "  0000000000012345678"  12345678
	// "00000-42a1234"  0
	// "21474836460"  2147483647
	// "+","+-","-+" 0
	flag := true
	res := 0
	ten := 10
	MAX_32 := (1 << 31) - 1
	MIN_32 := -(1 << 31)
	for i := 0; i < len(str); i++ {
		if str[i] == ' ' {
			continue
		} else if str[i] == '-' {
			flag = false
			if i+1 < len(str) && (str[i+1] < '0' || str[i+1] > '9') {
				return res
			}
		} else if str[i] == '+' {
			flag = true
			if i+1 < len(str) && (str[i+1] < '0' || str[i+1] > '9') {
				return res
			}
		} else if str[i] > '0' && str[i] <= '9' {
			if res > 0 {
				res *= ten
			}
			res += int(str[i] - '0')
		} else if str[i] == '0' {
			if res > 0 {
				res *= ten
			}
			res += int(str[i] - '0')
			break
		} else {
			break
		}
	}

	if res > 0 {
		if flag {
			if res >= MAX_32 {
				return MAX_32
			}
			return res
		} else {
			if res > MAX_32 {
				return MIN_32
			}
			return -res
		}
	}
	return 0
}

// 硬编码实现，正确的代码
func myAtoi(str string) int {
	sign := 1        // 正负数
	ans := 0         // 当前答案
	pop := 0         // 当前数字
	hasSign := false // 是否开始转换数字
	MAX := (1 << 31) - 1
	MIN := -(1 << 31)
	for i := 0; i < len(str); i++ {
		if str[i] == '-' && !hasSign {
			sign = -1
			hasSign = true
			continue
		}
		if str[i] == '+' && !hasSign {
			sign = 1
			hasSign = true
			continue
		}
		if str[i] == ' ' && !hasSign {
			continue
		}
		if str[i] >= '0' && str[i] <= '9' {
			hasSign = true
			pop = int(str[i] - '0')
			// 极值处理
			if ans*sign > MAX/10 || (ans*sign == MAX/10 && pop*sign > 7) {
				return MAX
			}
			if ans*sign < MIN/10 || (ans*sign == MIN/10 && pop*sign < -8) {
				return MIN
			}
			ans = ans*10 + pop
		} else {
			return ans * sign
		}
	}
	return ans * sign
}