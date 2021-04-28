package ArrayList

import "errors"

/*
	实现数组迭代器

	用指针访问数据(相当于ES6中的Proxy代理一样)
*/
type Iterator interface{
	HasNext() bool // 是否有下一个
	Next() (interface{},error) // 下一个
	Remove() // 删除
	GetIndex() int //得到索引
}

type Iterable interface{
	Iterator() Iterator // 构造初始化接口
}

// 构造指针访问数组
type ArrayListIterator struct{
	list *ArrayList // 数组指针
	currentIndex int // 当前索引
}

// 增加迭代器
func (list *ArrayList) Iterator() Iterator{
	// 构造迭代器
	it := new(ArrayListIterator)
	it.currentIndex = 0
	it.list = list
	return it
}

/*
	实现接口（todos：：这里不懂了）
*/
func(it *ArrayListIterator) HasNext() bool {
	return it.currentIndex < it.list.TheSize
}

func(it *ArrayListIterator) Next() (interface{},error) {
	if !it.HasNext(){
		return nil,errors.New("没有下一个")
	}
	// 获取当前数据
	value,err:=it.list.Get(it.currentIndex)
	it.currentIndex++
	return value,err
}

// 删除一个元素
func(it *ArrayListIterator) Remove() {
	it.currentIndex--
	it.list.Delete(it.currentIndex)
}

func(it *ArrayListIterator) GetIndex() int {
	return it.currentIndex
}