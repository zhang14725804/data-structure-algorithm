package CircleQueue

import "errors"

/*
	循环队列(一个环)
*/
const QUEUE_SIZE = 100

type CircleQueue struct{
	data [QUEUE_SIZE]interface{} // 存储数据的结构
	head int // 头部的位置
	tail int // 尾部的位置
}

// 初始化，头部尾部重合，为空
func InitQueue(q *CircleQueue)  {
	q.head = 0
	q.tail = 0
}

// todos::注意这里
func QueueLength(q *CircleQueue) int  {
	return (q.tail - q.head + QUEUE_SIZE) % QUEUE_SIZE
}

func EnQueue(q *CircleQueue, data interface{}) (err error)  {
	if (q.tail+1)%QUEUE_SIZE == (q.head)%QUEUE_SIZE{
		return errors.New("队列已满")
	}
	q.data[q.tail]= data // 入队
	q.tail = (q.tail+1)%QUEUE_SIZE
	return nil
}

// 出队
func DeQueue(q *CircleQueue) (data interface{},err error)  {
	if q.tail == q.head{
		return nil ,errors.New("队列为空")
	}
	res:=q.data[q.head] // 去除第一个数据
	q.data[q.head] = 0 // 清空数据
	q.head = (q.head+1) % QUEUE_SIZE // 注意这里
	return res,nil
}
