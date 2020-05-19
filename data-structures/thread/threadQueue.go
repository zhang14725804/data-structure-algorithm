// 线程安全的队列
type Queue struct{
	queue []interface{}
	Length int
	lock sync.Mutex // 安全锁
}

func New() *Queue  {
	q := &Queue{}
	q.queue = make([]interface{},0)
	q.Length = 0
	q.lock = new(sync.Mutex)
	return q
}


func (q *Queue) Len() int {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.Length
}