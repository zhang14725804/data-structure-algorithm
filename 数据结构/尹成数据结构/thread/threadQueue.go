// 线程安全的队列
type Queue struct{
	queue []interface{}
	Length int
	lock *sync.Mutex // 安全锁
}

func NewQueue() *Queue  {
	q := &Queue{}
	q.queue = make([]interface{},0)
	q.Length = 0
	q.lock = nil
	return q
}


func (q *Queue) Len() int {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.Length
}

func (q *Queue) Push(ele interface{}) {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.Length++
	q.queue = append(q.queue,ele)
}

func (q *Queue) Shift() interface{} {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.Length--
	res:=q.queue[0]
	q.queue = q.queue[1:]
	return res
}