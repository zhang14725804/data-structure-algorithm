/*
	给定两个由一些 闭区间 组成的列表，每个区间列表都是成对不相交的，并且已经排序。
	返回这两个区间列表的交集。

	（形式上，闭区间 [a, b]（其中 a <= b）表示实数 x 的集合，而 a <= x <= b。两个闭区间的交集是一组实数，要么为空集，要么为闭区间。例如，[1, 3] 和 [2, 4] 的交集为 [2, 3]。）
*/

/*
	解题思路：排序，画图，找规律！！！

	（1）不存在交集的情况两种
	（2）存在交集的情况四种
	交集区间是有规律的。对于[a1,a2],[b1,b2]两个区间，如果交集区间是 [c1,c2]，那么 c1=max(a1,b1)，c2=min(a2,b2)
	对于是否前进只取决于a2和b2的大小

	先找规律再做题，直接if else懵了😅
*/
func intervalIntersection(A [][]int, B [][]int) [][]int {
	i, j := 0, 0
	n, m := len(A), len(B)
	res := make([][]int, 0)
	for i < n && j < m {
		cur := make([]int, 2)
		a1, a2 := A[i][0], A[i][1]
		b1, b2 := B[j][0], B[j][1]
		// 相交的情况
		if a1 <= b2 && a2 >= b1 {
			cur[0] = MaxInt(a1, b1)
			cur[1] = MinInt(b2, a2)
			res = append(res, cur)
		}
		// 指针前进的条件
		if b2 < a2 {
			j++
		} else {
			i++
		}

	}
	return res
}