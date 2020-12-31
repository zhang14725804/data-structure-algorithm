package main

import "fmt"

/*
	堆的两个特性：
	结构性：用【数组】表示的【完全二叉树】
	有序性：大顶堆，小顶堆
*/

//
type Heap interface {
	IsFull() bool
	IsEmpty() bool
	Insert(it int)
	Delete() int
	BuildHeap(items []int) // 建立大顶堆（线性时间复杂度）
}

type heap struct {
	Element []int // 元素其实可以是任何类型，这里先用int类型。下标从1开始，设置一个哨兵
	// Min      bool  // 标记大根堆，或小根堆
	Size     int // 当前大小
	Capacity int // 容量
}

func NewHeap(cap int) Heap {
	elems := make([]int, 0)
	// 设置一个哨兵
	elems = append(elems, (1 << 32))
	return &heap{
		Element: elems,
		// Min:      min,
		Size:     0,
		Capacity: cap,
	}
}

func (h *heap) Insert(it int) {
	if h.IsFull() {
		fmt.Println("满了！！")
		return
	}

	h.Size++
	// 当前放置it的位置
	i := h.Size
	// 因为有哨兵，所以不需要判断i是否小于1
	for ; h.Element[i/2] < it; i /= 2 {
		h.Element[i] = h.Element[i/2]
	}
	h.Element[i] = it
}

func (h *heap) Delete() int {
	if h.IsEmpty() {
		fmt.Println("空了！！")
		return -1
	}
	max := h.Element[1]       // 最出最大元素，需要删除的元素
	temp := h.Element[h.Size] // 取出最后一个元素
	h.Size--
	var parent, child int
	// 寻找找插入点，判断是否有左节点；parent = child 进入下一层
	for parent = 1; parent*2 <= h.Size; parent = child {
		// 左节点
		child = parent * 2
		// child != h.Size 一定有右节点；比较左右节点 h.Element[child] < h.Element[child+1]
		if child != h.Size && h.Element[child] < h.Element[child+1] {
			child++
		}
		// 比左右节点都大
		if temp >= h.Element[child] {
			break
		} else {
			h.Element[parent] = h.Element[child]
		}
	}
	h.Element[parent] = temp
	return max
}

func (h *heap) IsEmpty() bool {
	return len(h.Element) == 0
}

func (h *heap) IsFull() bool {
	return h.Size == h.Capacity
}

/*
	方法1：一个个Insert到初始为空的堆中 (Nlog(N))
	方法2：先按照完全二叉树的结构特性将元素插入，然后由底向上调整堆（线性时间复杂度）
*/
func (h *heap) BuildHeap(items []int) {
	// 第一步
	for i := 0; i < len(items); i++ {
		h.Element = append(h.Element, items[i])
		h.Size++
	}
	// （ TODO ）第二步，不会呀😅
}
