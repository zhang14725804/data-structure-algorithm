/*
	输入一个正数s，打印出所有和为s的连续正数序列（至少含有两个数）
	todo：双指针算法，思路清奇。怎么想出来的😅，思路不懂，有点难
*/
func findContinuousSequence(sum int) [][]int {
	res := make([][]int,0)
	for i,j,s:=1,1,1;i<=sum;i++{
		for s < sum{
			j++
			s+=j
		}
		if s == sum && j-i>0{
			line := make([]int,0)
			for k:=i;k<=j;k++{
				line = append(line,k)
			}
			res = append(res,line)
		}
		s-=i
	}
	return res
}
