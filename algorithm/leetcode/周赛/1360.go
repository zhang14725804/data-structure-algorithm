/*
	思路：前缀和
	1971.1.1-->day1
	1971.1.1-->day2
	年、月、日分开算
*/
func daysBetweenDates(date1 string, date2 string) int {
	return int(math.Abs(float64(getDay(date1)-getDay(date2))))
}

// 是否是闰年（两种情况）
func isLeap(year int) int {
	if year%100 !=0 && year %4==0 || year%400==0{
		return 1
	}
	return 0
}
var monthDays = []int{0,31,28,31,30,31,30,31,31,30,31,30,31}
func getDay(date string) int  {
	var year,month,day int
	// 从字符串取出年月日
	fmt.Sscanf(date,"%d-%d-%d",&year,&month,&day)
	days:=0

	// 年
	for i:=1971;i<year;i++{
		days+=365+isLeap(i)
	}
	// 月
	for i:=1;i<month;i++{
		if i==2{
			days+=monthDays[i]+isLeap(year)
		}else{
			days+=monthDays[i]
		}
	}
	// 日
	days+=day
	return days
}