package main

import "fmt"

type Set struct{
	buf []interface{} // 存储数据
	nums int // 数量
	hash map[interface{}]bool
}

func newSet() *Set  {	
	return &Set{
		buf: make([]interface{},0),
		nums: 0,
		hash: make(map[interface{}]bool),
	}
}

func (set *Set) Add(val interface{}) bool {
	if set.isExit(val){
		return false
	}else{
		set.buf = append(set.buf,val)
		set.hash[val] = true
		set.nums++
		return true
	}
}

func (set *Set) isExit(val interface{}) bool {
	return set.hash[val]
}

// 
func (set *Set) Strings() []interface{} {
	return set.buf
}
/*
	hash表原理（适合查询，删除修改增加不适合）
*/
func main()  {
	set := newSet()
	set.Add(1)
	set.Add(2)
	fmt.Println(set)
}