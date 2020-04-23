package gobase

import (
	"fmt"
	"reflect"
)

// Structure 结构体
type Structure struct{}

// ReflectStructMain 反射场景实践
/*
	参考文章
	https://segmentfault.com/a/1190000016230264
	https://juejin.im/post/5a75a4fb5188257a82110544
*/
func ReflectStructMain() {
	name := "CallbackWithoutParams"
	t := &Structure{}
	v := reflect.ValueOf(t)
	fmt.Println(v)
	k := v.Kind()
	fmt.Println(k)
	mbn := v.MethodByName((name))
	fmt.Println(mbn)
	mbn.Call(nil)
}

// CallbackWithoutParams 动态调用无餐函数
func (t *Structure) CallbackWithoutParams() {
	fmt.Println("=====CallbackWithoutParams=====")
}
