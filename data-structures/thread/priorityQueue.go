/*
	优先队列：一类抽象数据类型。
	优先队列中的每个元素都有各自的优先级，优先级最高的元素最先得到服务；优先级相同的元素按照其在优先队列中的顺序得到服务。优先队列往往用堆来实现（todo：为什么用堆来实现）。
*/ 

type PriorityItem struct{
	value interface{} // 数据
	priority int // 优先级
}

func NewPriorityItem(value interface{}, priority int) *PriorityItem  {
	return &PriorityItem{value, priority}
}

// 比较大小设定优先级
func (a PriorityItem) Less(b Item) bool {
	return a.priority < b.(PriorityItem).priority 
}

// 基于堆的优先队列
type PriorityQueue struct{
	data Heap
}

func NewMaxPriorityQueue() *PriorityQueue {
	return &PriorityQueue{*NewMax()}
}

func NewMinPriorityQueue() *PriorityQueue {
	return &PriorityQueue{*NewMin()}
}

func (pq *PriorityQueue) Len() int  {
	return pq.data.Len()
}

func (pq *PriorityQueue) Insert(ele PriorityItem)  {
	pq.data.Insert(ele)
}

// 
func (pq *PriorityQueue) Extract() PriorityItem {
	return pq.data.Extract().(PriorityItem)
}

// todo:干了些啥
func (pq *PriorityQueue) ChangePriority(value interface{}, priority int) {
	// 队列备份数据
	storage := NewQueue()
	p:=pq.Extract()
	if value !=p.value{
		if pq.Len() == 0{
			return
		}
		storage.Push(p)
		p=pq.Extract()
	}
	p.priority = priority
	pq.data.Insert(p)
	// 其余数据重新放入优先队列
	for storage.Len() > 0{
		pq.data.Insert(storage.Shift().(Item))
	}
}

func main()  {
	h := NewMinPriorityQueue()
	h.Insert(*NewPriorityItem(101,11))
	h.Insert(*NewPriorityItem(102,12))
	h.Insert(*NewPriorityItem(103,15))
	h.Insert(*NewPriorityItem(104,14))
	h.Insert(*NewPriorityItem(105,13))
	h.Insert(*NewPriorityItem(106,19))

	fmt.Println(h.Extract())
}