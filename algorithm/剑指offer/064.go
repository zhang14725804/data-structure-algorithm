/*
	找出字符流中第一个只出现一次的字符
	用63题方法，复杂度n*n。要优化成n
	todo：思路不错😄
*/

	
var q queue
var count = make(map[byte]int)
func insert(ch byte) {
	/*
		如何改造++count[ch] > 1
	*/
	count[ch]++
	if count[ch] > 1{
		for q.size() > 0 && count[q.front()] > 1{
			q.pop()
		}
	}else{
		q.x = append(q.x, ch)
	}
}
func firstAppearingOnce() byte {
	
	if q.size() == 0{
		return '#'
	}
	return q.front()
}




// 队列
type queue struct{
	x []byte
}

// 尾插入
func (this *queue) push(x byte){
	this.x = append(this.x,x)
}

// 返回头元素int
func (this *queue) front() byte{
	return this.x[0]
}
// size
func (this *queue) size() int{
	return len(this.x)
}
// 头删除
func (this *queue) pop(){
	this.x= this.x[1:]
}