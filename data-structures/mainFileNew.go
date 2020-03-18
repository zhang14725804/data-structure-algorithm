package main

import (
	"errors"
	"fmt"
	"io/ioutil"
)

// 递归遍历文件夹
func GetAllNew(path string,files []string)([]string,error)  {
	// 读取文件夹
	read,err:=ioutil.ReadDir(path)
	if err!=nil{
		return files,errors.New("文件夹不可读")
	}
	// 循环每个文件或者文件夹
	for _,fi := range read{
		fullDir := path+"\\"+fi.Name() // 构造新的路径
		files=append(files,fullDir) // 追加路径
		if fi.IsDir(){ // 判断是否是文件夹
			files,_=GetAllNew(fullDir,files) //文件夹递归处理
		}
	}
	return files,nil
}

/*
	递归的方式遍历文件夹
	todos：：显示文件层级level（目前层级有问题，只能遍历到第二层）
*/
func main()  {
	path := "D:\\aprivate"
	files:=[]string{}
	files,_=GetAllNew(path,files)
	// 循环打印文件或者文件夹
	for i:=0;i<len(files) ; i++  {
		fmt.Println(files[i])
	}
}