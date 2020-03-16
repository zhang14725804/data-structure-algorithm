package ArrayList

import (
	"errors"
	"fmt"
)

/*
	数组，内存排列在一起
*/

// 接口
type List interface {
	Size() int                                  // 数组大小
	Get(index int) (interface{}, error)         // 获取第几个元素
	Set(index int, newVal interface{}) error    // 修改数据
	Insert(index int, newVal interface{}) error // 插入数据
	Append(newVal interface{})                  // 追加
	Clear()                                     // 清空
	Delete(index int) error                     // 删除
	String() string                             // 返回字符串
	Iterator() Iterator
}

// 数据结构，字符串，整数，实数(interface{}任意类型)
type ArrayList struct {
	dataStore []interface{} // 数组存储
	TheSize   int           // 数组大小(注意：theSize首字母小写表示私有类型)
}

func NewArrayList() *ArrayList {
	list := new(ArrayList)                      // 初始化结构体
	list.dataStore = make([]interface{}, 0, 10) // 开辟空间10个
	list.TheSize = 0
	return list
}

func (list *ArrayList) Size() int {
	return list.TheSize // 返回数组大小
}

func (list *ArrayList) Append(newVal interface{}) {
	list.dataStore = append(list.dataStore, newVal)
	list.TheSize++
}

// 获取数据
func (list *ArrayList) Get(index int) (interface{}, error) {
	if index < 0 || index >= list.TheSize {
		return nil, errors.New("索引越界")
	}
	return list.dataStore[index], nil
}

func (list *ArrayList) String() string {
	return fmt.Sprint(list.dataStore)
}

// 设置数据
func (list *ArrayList) Set(index int, newVal interface{}) error {
	if index < 0 || index >= list.TheSize {
		return errors.New("索引越界")
	}
	list.dataStore[index] = newVal
	return nil
}

/*
	插入数据
	👇
*/
func (list *ArrayList) Insert(index int, newVal interface{}) error {
	if index < 0 || index >= list.TheSize {
		return errors.New("索引越界")
	}
	list.checkIsFull()
	// 插入数据，内存移动一位
	list.dataStore = list.dataStore[:list.TheSize+1]
	// 从后往前移动
	for i := list.TheSize; i > index; i-- {
		list.dataStore[i] = list.dataStore[i-1]
	}
	list.dataStore[index] = newVal
	list.TheSize++
	return nil
}

/*
	查看数组是否满(这里要注意)
	👇
*/
func (list *ArrayList) checkIsFull() {
	if list.TheSize == cap(list.dataStore) {
		// 开辟双倍内存（这里要注意）
		newDataStore := make([]interface{}, 2*list.TheSize, 2*list.TheSize)
		/*
			copy等价于这个
			copy(newDataStore,list.dataStore)
		*/
		for i:=0;i<len(list.dataStore)-1;i++{
			newDataStore[i] = list.dataStore[i]
		}
		// 赋值
		list.dataStore = newDataStore
	}
}

// 清空数组
func (list *ArrayList) Clear() {
	list.dataStore = make([]interface{}, 0, 10) // 开辟空间10个
	list.TheSize = 0
}

/*
	删除数据（注意展开运算符）
	👇
*/
func (list *ArrayList) Delete(index int) error {
	list.dataStore = append(list.dataStore[:index], list.dataStore[index+1:]...)
	list.TheSize--
	return nil
}


