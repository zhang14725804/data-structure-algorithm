
type Int int

func (x Int) Less(than Item) bool {
	return x<than.(Int)
}

// 接口实现对比大小
type Item interface{
	Less(than Item) bool
}

// 最大堆、最小堆
type Heap struct{
	lock *sync.Mutex // 线程锁
	data []Item
	min bool // 是否是最小堆
}

// 标准堆，最大堆，最小堆
func NewHeap() *Heap  {
	return &Heap{new(sync.Mutex),make([]Item,0),true}
}

func NewMin() *Heap {
	return &Heap{new(sync.Mutex),make([]Item,0),true}
}

func NewMax() *Heap {
	return &Heap{new(sync.Mutex),make([]Item,0),false}
}

func (h *Heap)isEmpty() bool {
	return len(h.data) == 0
}

func (h *Heap) Len() int {
	return len(h.data)
}

// 获取数据
func (h *Heap) Get(index int) Item {
	return h.data[index]
}

func (h *Heap) Insert(it Item)  {
	//h.lock.Lock()
	//defer h.lock.Unlock()
	h.data = append(h.data,it)
	// todo:这是要干嘛
	h.shiftUp()
}

// 根据堆类型，返回比较大小
func (h *Heap) Less(a,b Item) bool {
	if h.min{
		return a.Less(b)
	}
	return b.Less(a)
}

// 调整堆
func (h *Heap) Extract() Item {
	// h.lock.Lock()
	// defer h.lock.Unlock()
	if h.Len() == 0{
		return nil 
	}
	// 第一个元素和最后一个元素
	el:=h.data[0]
	last:=h.data[h.Len()-1]
	if h.Len()==1{
		h.data = nil // 重新分配内存
		return nil
	}
	// todo:没看懂
	h.data = append([]Item{last},h.data[1:h.Len()-1]...)
	// todo:这是要干嘛
	h.shiftDown()
	return el // 返回第一个元素
}

// 弹出极大值极小值(todo:没看懂)
func (h *Heap) shiftUp(){
	// 堆排序
	for i,parent:=h.Len()-1,h.Len()-1;i>0;i=parent{
		parent = i/2
		if h.Less(h.Get(i),h.Get(parent)){
			h.data[parent],h.data[i]=h.data[i],h.data[parent]
		}else{
			break
		}
	}
} 
func (h *Heap) shiftDown(){
	for i,child:=0,1;i<h.Len() && i*2+1<h.Len();i=child{
		child = i*2 +1
		if child+1 <= h.Len()-1 && h.Less(h.Get(child+1),h.Get(child)){
			child++ 
		}
		if h.Less(h.Get(i),h.Get(child)){
			break
		}
		h.data[i],h.data[child]=h.data[child],h.data[i]
	}
} 

func main(){
	h := NewMin()
	h.Insert(Int(8))
	h.Insert(Int(9))
	h.Insert(Int(7))
	h.Insert(Int(5))
	h.Insert(Int(6))
	h.Insert(Int(4))
	h.Insert(Int(2))
	h.Insert(Int(3))

	fmt.Println(h.Extract().(Int))
}