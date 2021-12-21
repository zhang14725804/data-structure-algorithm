/*
	本题有两个维度，h和k，看到这种题目一定要想如何确定一个维度，然后在按照另一个维度重新排列。

	【遇到两个维度权衡的时候，一定要先确定一个维度，再确定另一个维度。】

	对于本题相信大家困惑的点是先确定k还是先确定h呢，也就是究竟先按h排序呢，还先按照k排序呢？

	如果按照k来从小到大排序，排完之后，会发现k的排列并不符合条件，身高也不符合条件，两个维度哪一个都没确定下来。

	那么按照身高h来排序呢，身高一定是从大到小排（身高相同的话则k小的站前面），让高个子在前面。

	此时我们可以确定一个维度了，就是身高，前面的节点一定都比本节点高！

	那么只需要按照k为下标重新插入队列就可以了，为什么呢？

	按照身高排序之后，优先按身高高的people的k来插入，后序插入节点也不会影响前面已经插入的节点，最终按照k的规则完成了队列。

	所以在按照身高从大到小排序后：

	局部最优：优先按身高高的people的k来插入。插入操作过后的people满足队列属性

	全局最优：最后都做完插入操作，整个队列满足题目队列属性

*/

func reconstructQueue(people [][]int) [][]int {
	people = BubbleSort(people)
	pLen := len(people)
	var res [][]int

	for i := 0; i < pLen; i++ {
		idx := people[i][1]
		// 😅😅😅
		// c++这么写： que.insert(que.begin() + position, people[i]);
		res = append(res[:idx], append([][]int{people[i]}, res[idx:]...)...)
	}
	return res
}

/*
	身高高的排前面，身高相同的序号小的排前面
*/
func compare(a, b []int) bool {
	if a[0] == b[0] {
		return a[1] < b[1]
	}
	return a[0] > b[0]
}