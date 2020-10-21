/*
	给定一个迭代器类的接口，接口包含两个方法： next() 和 hasNext()。设计并实现一个支持 peek() 操作的顶端迭代器 -- 其本质就是把原本应由 next() 方法返回的元素 peek() 出来。

	你将如何拓展你的设计？使之变得通用化，从而适应所有的类型，而不只是整数型？
*/
/*   Below is the interface for Iterator, which is already defined for you.
 *
 *   type Iterator struct {
 *
 *   }
 *
 *   func (this *Iterator) hasNext() bool {
 *		// Returns true if the iteration has more elements.
 *   }
 *
 *   func (this *Iterator) next() int {
 *		// Returns the next element in the iteration.
 *   }
 */
type PeekingIterator struct {
}

func Constructor(iter *Iterator) *PeekingIterator {

}

func (this *PeekingIterator) hasNext() bool {

}

func (this *PeekingIterator) next() int {

}

func (this *PeekingIterator) peek() int {

}