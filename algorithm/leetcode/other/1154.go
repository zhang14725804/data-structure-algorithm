/*
   1. 拆分字符串，得到年、月、日
   2. 根据年份判断平年还是闰年，进而决定2月有几天
   3. 根据月份判断这个月有多少天
   4. 月日相加得到最终结果
*/
func dayOfYear(date string) int {
	ymdArr := strings.Split(date, "-")
	year := str2int(ymdArr[0])
	month := str2int(ymdArr[1])
	day := str2int(ymdArr[2])

	ds := 0
	for month > 0 {
		ds += getMonthDays(month, year)
		month--
	}

	fmt.Println(ds + day)

	return ds + day
}

func str2int(str string) int {
	ten := 10
	bstr := []byte(str)
	res := 0
	for i := 0; i < len(bstr); i++ {
		res = res*ten + int(bstr[i]-'0')
	}
	return res
}

// 判断闰年平年，1. 4的倍数，且不是100的倍数 2. 整百数的，必须是400的倍数
func isRun(year int) bool {
	if year%100 == 0 {
		return year%400 == 0
	}
	if year%4 == 0 {
		return year%100 > 0
	}
	return false
}

func getMonthDays(month, year int) int {
	month -= 1
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		return 31
	case 4, 6, 9, 11:
		return 30
	case 2:
		if isRun(year) {
			return 29
		}
		return 28
	}
	return 0
}
