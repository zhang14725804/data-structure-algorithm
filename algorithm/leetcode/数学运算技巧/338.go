/*
	给定一个非负整数 num。对于 0 ≤ i ≤ num 范围中的每个数字 i ，计算其二进制数中的 1 的数目并将它们作为数组返回。

	给出时间复杂度为O(n*sizeof(integer))的解答非常容易。但你可以在线性时间O(n)内用一趟扫描做到吗？
	要求算法的空间复杂度为O(n)。
	你能进一步完善解法吗？要求在C++或任何其他语言中不使用任何内置函数（如 C++ 中的 __builtin_popcount）来执行此操作。
*/

/*
	方法1：用191题的思路
*/
func countBits(num int) []int {
	// 从0开始
	res := make([]int, num+1)
	for i := 0; i <= num; i++ {
		res[i] = count(i)
	}
	return res
}

func count(n int) (res int) {
	for n > 0 {
		// n&(n-1) 😅😅😅 ：消除数字 n 的二进制表示中的最后一个 1
		n = n & (n - 1)
		res++
	}
	return
}

/*
	方法2：动态规划 question 😅😅😅
*/
func countBits(num int) []int {
	res := make([]int, num+1)
	for i := 1; i <= num; i++ {
		// 上一个缺1的元素+1即可
		res[i] = res[i&(i-1)] + 1
	}
	return res
}