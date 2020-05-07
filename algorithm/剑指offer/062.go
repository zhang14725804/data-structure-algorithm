/*
	我们把只包含质因子2、3和5的数称作丑数（Ugly Number）
	求第n个丑数的值
	todo：思路：这思路有点绕😅,不懂
*/
func getUglyNumber(n int) int {
	q:=make([]int,1)
	q[0] = 1
	i,j,k := 0,0,0
	n = n-1
	for n > 0{
	    n--
	    t := min(q[i]*2, min(q[j]*3, q[k]*5))
		q = append(q,t)
		if q[i]*2 == t {
			i++
		}
		if q[j]*3 == t {
			j++
		}
		if q[k]*5 == t {
			k++
		}
	}
	return q[len(q)-1]
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}