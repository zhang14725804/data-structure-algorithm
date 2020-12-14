/*  
    给你一个嵌套的整型列表。请你设计一个迭代器，使其能够遍历这个整型列表中的所有整数。
    列表中的每一项或者为一个整数，或者是另一个列表。其中列表的元素也可能是整数或是其他列表。

    迭代器惰性求值（每次只循环获取下一个元素）
    N叉树遍历问题，dfs或者bfs会一次性把所有结果都算出来
    迭代器惰性求值：调用hasNext时，如果nestedList的第一个元素是列表类型，则不断展开这个元素，直到第一个元素是整数类型
*/
 
type NestedIterator struct {
	seq []int
	cnt int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
    dfs(nestedList)
}


func (this *NestedIterator) Next() int {
    return this.seq[this.cnt++]
}

func (this *NestedIterator) HasNext() bool {
    return this.cnt == len(this.seq)
}