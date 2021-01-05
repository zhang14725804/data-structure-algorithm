/*
	给定一个范围在  1 ≤ a[i] ≤ n ( n = 数组大小 ) 的 整型数组，数组中的元素一些出现了【两次】，另一些只出现一次。
	找到所有在 [1, n] 范围之间没有出现在数组中的数字。
	您能在不使用额外空间且时间复杂度为O(n)的情况下完成这个任务吗? 你可以假定返回的数组不算在额外空间内。

	如何寻找缺失的元素（ question 🔥🔥🔥 ）
	对于缺失一个数而且不存在重复数字的情况：
	方法1：排序然后查找
	方法2：借用hashSet
	方法3：位运算，对于异或运算（^）：一个数和它本身做异或运算结果为 0，一个数和 0 做异或运算还是它本身
	方法4：等差数列求和，然后再减去当前数组中的数，注意【整型溢出 ❗】
*/
func findDisappearedNumbers(nums []int) []int {
	n := len(nums)
	used := make([]bool, n)
	for i := 0; i < n; i++ {
		used[nums[i]-1] = true
	}
	res := make([]int, 0)
	for i := 0; i < n; i++ {
		if !used[i] {
			res = append(res, i+1)
		}
	}
	return res
}