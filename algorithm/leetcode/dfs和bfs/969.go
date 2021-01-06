/*
	给定数组 A，我们可以对其进行煎饼翻转：我们选择一些正整数 k <= A.length，然后反转 A 的前 k 个元素的顺序。我们要执行零次或多次煎饼翻转（按顺序一次接一次地进行）以完成对数组 A 的排序。
	返回能使 A 排序的煎饼翻转操作所对应的 k 值序列。任何将数组排序且翻转次数在 10 * A.length 范围内的有效答案都将被判断为正确。
*/

/*
	方法1：递归
	1、找到n个饼中最大的那个。
	2、把这个最大的饼移到最底下。
	3、递归调用pancakeSort(A, n - 1)。
	base case：n == 1时，排序 1 个饼时不需要翻转。
	【如何设法将某块烧饼翻到最后呢？】
	比如第 3 块饼是最大的，我们想把它换到最后，也就是换到第n块。可以这样操作：
	1、用锅铲将前 3 块饼翻转一下，这样最大的饼就翻到了最上面。
	2、用锅铲将前n块饼全部翻转，这样最大的饼就翻到了第n块，也就是最后一块。
*/

var res []int

func pancakeSort(arr []int) []int {
	dfs(arr, len(arr))
	return res
}

// 可惜不是最优解 😅
// question 计算排序烧饼的【最短】操作序列，应该是要用BFS（不对，回溯求出所有可行的解然后找到最优解），问题是怎么写😅
func dfs(cakes []int, n int) {
	// base case
	if n == 1 {
		return
	}
	// 寻找最大的饼索引
	maxCake := 0
	maxCakeIndex := 0
	for i := 0; i < len(cakes); i++ {
		if cakes[i] > maxCake {
			maxCake = cakes[i]
			maxCakeIndex = i
		}
	}
	// 第一次反转
	reverse(cakes, 0, maxCakeIndex)
	res = append(res, maxCakeIndex+1)
	// 第二次反转
	reverse(cakes, 0, n-1)
	res = append(res, n)
	// 递归调用
	dfs(cakes, n-1)
}
func reverse(cakes []int, i, j int) {
	for i < j {
		cakes[i], cakes[j] = cakes[j], cakes[i]
		i++
		j--
	}
}
	