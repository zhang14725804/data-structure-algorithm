/*
	整数转罗马数字
*/

// hash := map[string]int{
// 	"I":  1,
// 	"IV": 4,
// 	"V":  5,
// 	"IX": 9,
// 	"X":  10,
// 	"XL": 40,
// 	"L":  50,
// 	"XC": 90,
// 	"C":  100,
// 	"CD": 400,
// 	"D":  500,
// 	"CM": 900,
// 	"M":  1000,
// }
func intToRoman2(num int) string {
	// 是有多蠢😅
	res := ""
	for num > 0 {
		if num/1000 > 0 {
			n := num / 1000
			for n > 0 {
				res += "M"
				n--
			}
			num %= 1000
		} else if num-900 >= 0 {
			res += "CM"
			num -= 900
		} else if num-500 >= 0 {
			res += "D"
			num -= 500
		} else if num-400 >= 0 {
			res += "CD"
			num -= 400
		} else if num-100 >= 0 {
			res += "C"
			num -= 100
		} else if num-90 >= 0 {
			res += "XC"
			num -= 90
		} else if num-50 >= 0 {
			res += "L"
			num -= 50
		} else if num-40 >= 0 {
			res += "XL"
			num -= 40
		} else if num-10 >= 0 {
			res += "X"
			num -= 10
		} else if num-9 >= 0 {
			res += "IX"
			num -= 9
		} else if num-5 >= 0 {
			res += "V"
			num -= 5
		} else if num-4 >= 0 {
			res += "IV"
			num -= 4
		} else {
			res += "I"
			num -= 1
		}
	}
	return res
}

func intToRoman(num int) string {
	thousands := []string{"", "M", "MM", "MMM"}
	hundreds := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	tens := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	ones := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}

	return thousands[num/1000] + hundreds[num%1000/100] + tens[num%100/10] + ones[num%10]
}