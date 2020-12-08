/*
	1390
*/
func sumFourDivisors(nums []int) int {
	res:=0
	// 遍历所有的数
	for _,x:=range nums{
		// 约数总和
		sum:=0
		// 约数个数
		cnt:=0
		// 枚举
		for i:=1;i*i<=x;i++{
			if x%i == 0{
				sum+=i
				cnt++
				// 排除约数相等的情况 6*6 = 36
				if x/i != i{
					sum+=x/i
					cnt++
				}
			}
		}
		if cnt==4{
			res+=sum
		}
	}
	return res
}