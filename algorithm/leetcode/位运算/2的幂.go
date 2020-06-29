/*
	给定一个整数，编写一个函数来判断它是否是 2 的幂次方。
	对于N为2的幂的数，都有 N&(N-1)=0 ，所以这就是我们的判断条件。
*/
func isPowerOfTwo(n int) bool {
	return n > 0 && n&(n-1)==0
}