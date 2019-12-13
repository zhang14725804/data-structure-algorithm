package gobase

import (
	"fmt"
	"reflect"
)

func reflectType(x interface{}) {
	obj := reflect.TypeOf(x)
	fmt.Println(obj, obj.Name(), obj.Kind())
	// *reflect.rtype 指针类型
	// fmt.Printf("format:%T\n", obj)
}

func reflectValue(x interface{}) {
	v := reflect.ValueOf((x))
	fmt.Printf("format:%v,%T\n", v, v)
	k := v.Kind()
	// 如何得到传入时的类型变量
	switch k {
	case reflect.Float32:
		ret := float32(v.Float())
		fmt.Printf("format:%v,%T\n", ret, ret)
	}
}

func reflectSetValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Elem().Kind()
	switch k {
	case reflect.Int32:
		v.Elem().SetInt(100)
	case reflect.Float32:
		v.Elem().SetFloat(3.14)
	}
}

type Dog struct{}
type Cat struct{}

// ReflectMain 主方法
func ReflectMain() {
	fmt.Println("====reflectType=====")
	var a float32 = 1.234
	reflectType(a)
	var b int64 = 33
	reflectType(b)

	var d Dog
	reflectType(d)

	var e []int
	reflectType(e)
	var f []string
	reflectType(f)

	var aaa float32 = 2.34
	fmt.Println("====reflectValue=====")
	reflectValue(aaa)
	reflectSetValue(&aaa)
	fmt.Println("====reflectSetValue=====")
	fmt.Println(aaa)
}
