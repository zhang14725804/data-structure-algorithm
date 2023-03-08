/*
	英文数字读法，三个三个分一组 😅😅😅
	1. 枚举0～19
	2. 枚举十位，但是没有0和10
	3. 枚举千分位，而且是倒叙，没有0
	4. 先算千分位
	5. 再算百位；十位(>=20)；20以内
*/
// 0-19
var small = [20]string{"Zero", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten",
	"Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}

// 十位（0和1没有）
var decade = [10]string{"", "", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}

// 1000（每三位一组）
var big = [4]string{"Billion", "Million", "Thousand", ""}

func numberToWords(num int) string {
	if num == 0 {
		return small[0]
	}

	res := ""
	// i数字，j单位
	for i, j := 1000000000, 0; i > 0; i, j = i/1000, j+1 {
		if num >= i {
			res += getPart(num/i) + big[j] + " "
			// 取余
			num %= i
		}
	}
	return strings.TrimSpace(res)
}
func getPart(num int) string {
	res := ""
	// 百位
	if num >= 100 {
		res += small[num/100] + " Hundred "
		// 取余数
		num %= 100
	}

	if num == 0 {
		return res
	}

	// 十位
	if num >= 20 {
		res += decade[num/10] + " "
		// 取余数
		num %= 10
	}

	if num == 0 {
		return res
	}

	// 个位
	res += small[num] + " "

	return res
}