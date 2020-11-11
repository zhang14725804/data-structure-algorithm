/*
	请你来实现一个 atoi 函数，使其能将字符串转换成整数。

	首先，该函数会根据需要丢弃无用的开头空格字符，直到寻找到第一个非空格的字符为止。接下来的转化规则如下：

	如果第一个非空字符为正或者负号时，则将该符号与之后面尽可能多的连续数字字符组合起来，形成一个有符号整数。
	假如第一个非空字符是数字，则直接将其与之后连续的数字字符组合起来，形成一个整数。
	该字符串在有效的整数部分之后也可能会存在多余的字符，那么这些字符可以被忽略，它们对函数不应该造成影响。
	注意：假如该字符串中的第一个非空格字符不是一个有效整数字符、字符串为空或字符串仅包含空白字符时，则你的函数不需要进行转换，即无法进行有效转换。

	在任何情况下，若函数不能进行有效的转换时，请返回 0 。

	提示：
	本题中的空白字符只包括空格字符 ' ' 。
	假设我们的环境只能存储 32 位大小的有符号整数，那么其数值范围为 [−231,  231 − 1]。如果数值超过这个范围，请返回  INT_MAX (231 − 1) 或 INT_MIN (−231) 。
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

type Automaton struct {
	Sign  int
	Ans   int
	state string
	table map[string][]string
}

func (auto Automaton) get(c byte) {
	auto.state = auto.table[auto.state][auto.getCol(c)]
	if "in_number" == auto.state {
		auto.Ans = auto.Ans*10 + int(c-'0')
		if auto.Sign == 1 {
			auto.Ans = MinInt((1<<31)-1, auto.Ans)
		} else {
			auto.Ans = MaxInt(-(1 << 31), -auto.Ans)
		}
	} else if "signed" == auto.state {
		if c == '+' {
			auto.Sign = 1
		} else {
			auto.Sign = -1
		}
	}
}

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

// 直接上手写程序的结果，还没有写出来😅
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