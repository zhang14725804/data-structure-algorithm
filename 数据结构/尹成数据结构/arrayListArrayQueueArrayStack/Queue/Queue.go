package Queue

/*
	注意啦（todos）：：
		队列是广度优先遍历！！！
		栈是深度优先遍历（递归也是）
*/
type MyQueue interface{
	Size() int // 大小
	Front() interface{} // 第一个元素
	End() interface{} //最后一个元素
	IsEmpty() bool // 是否为空
	EnQueue(data interface{}) // 入队
	DeQueue() interface{} // 出队
	Clear() //清空
}

type Queue struct{
	dataStore []interface{} // 队列的数据
	TheSize int // 大小
}

func NewQueue() *Queue{
	q := new(Queue) // 初始化，开辟结构体
	q.dataStore = make([]interface{},0)
	q.TheSize = 0
	return q
}

func (q *Queue) Size() int{
	return q.TheSize
}

func (q *Queue) Front() interface{} {
	if q.TheSize == 0{
		return nil
	}
	return q.dataStore[0]
}

func (q *Queue) End() interface{} {
	if q.TheSize ==0 {
		return nil
	}
	return q.dataStore[q.Size()-1]
}

func (q *Queue) IsEmpty() bool {
	return q.TheSize == 0
}

func (q *Queue) EnQueue(data interface{}) {
	q.dataStore = append(q.dataStore,data)
	q.TheSize++
}

/*
	如何实现循环队列（当前是顺序队列）
*/
func (q *Queue) DeQueue() interface{} {
	if q.Size() == 0{
		return nil
	}
	data := q.dataStore[0]
	if q.Size() > 1{
		q.dataStore = q.dataStore[1:q.Size()]
	}else{
		q.dataStore = make([]interface{},0)
	}
	q.TheSize--
	return data
}

func (q *Queue) Clear(){
		q.dataStore = make([]interface{},0)
		q.TheSize = 0
}