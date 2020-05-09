/*
	给定一个数组A[0, 1, …, n-1]，请构建一个数组B[0, 1, …, n-1]，其中B中的元素B[i]=A[0]×A[1]×… ×A[i-1]×A[i+1]×…×A[n-1]。

	不能使用除法。只能使用一个数组
	todo：稀里糊涂的没懂😅
	输入：[1, 2, 3, 4, 5]
	输出：[120, 60, 40, 30, 24]
*/
func multiply(a []int) []int {
	n := len(a)
	b := make([]int,n)
    if n == 0{
		return b
	}
	// 先算左边
	for i,p:=0,1; i < n; i++ {
		b[i] = p
		p *= a[i]
	}
	// 再算右边
	for i,p:=n-1,1;i>=0;i--{
		b[i] *= p
		p *= a[i]
	}
	return b
}