 

import (
	"./Queue"
	"./StackArray"
	"errors"
	"fmt"
	"io/ioutil"
)

// 递归遍历文件夹
func GetAll(path string, files []string) ([]string, error) {
	// 读取文件夹
	read, err := ioutil.ReadDir(path)
	if err != nil {
		return files, errors.New("文件夹不可读")
	}
	// 循环每个文件或者文件夹
	for _, fi := range read {
		fullDir := path + "\\" + fi.Name() // 构造新的路径
		files = append(files, fullDir)     // 追加路径
		if fi.IsDir() {                    // 判断是否是文件夹
			files, _ = GetAll(fullDir, files) //文件夹递归处理
		}
	}
	return files, nil
}

// 递归的方式遍历文件夹
func main1x() {
	path := "D:\\aprivate"
	files := []string{}
	files, _ = GetAll(path, files)
	// 循环打印文件或者文件夹
	for i := 0; i < len(files); i++ {
		fmt.Println(files[i])
	}
}

// 栈的方式遍历文件夹
func main2x() {
	path := "D:\\aprivate"
	files := []string{}
	stack := StackArray.NewStack()
	stack.Push(path)
	for !stack.IsEmpty() {
		// todos注意：这里这个变量名有点坑，这里的path和上面的path名相同，如果不一致了，递归不完整(为啥！！！)
		path := stack.Pop().(string)    // 类型转换
		files = append(files, path)     // 加入列表
		read, _ := ioutil.ReadDir(path) // 读取文件夹下面的路径
		for _, fi := range read {
			fullDir := path + "\\" + fi.Name() // 构造新的路径
			files = append(files, fullDir)     // 追加路径
			if fi.IsDir() {
				stack.Push(fullDir)
			}
		}
	}
	// 循环打印文件或者文件夹
	for i := 0; i < len(files); i++ {
		fmt.Println(files[i])
	}
}

// 用队列实现文件夹遍历
func main() {
	path := "D:\\aprivate"
	files := []string{}

	q := Queue.NewQueue()
	q.EnQueue(path)
	// 死循环
	for {
		path := q.DeQueue() // 类型转换
		if path == nil {
			break
		}
		fmt.Println("get ", path)
		read, _ := ioutil.ReadDir(path.(string)) // 读取文件夹下面的路径
		for _, fi := range read {
			// 文件夹和文件分别处理
			if fi.IsDir() {
				fullDir := path.(string) + "\\" + fi.Name() // 构造新的路径
				fmt.Println("Dir ", fullDir)
				q.EnQueue(fullDir)
			} else {
				fullDir := path.(string) + "\\" + fi.Name() // 构造新的路径
				files = append(files, fullDir)              // 追加路径
				fmt.Println("File ", fullDir)
			}
		}
	}
	// 循环打印文件或者文件夹
	for i := 0; i < len(files); i++ {
		fmt.Println(files[i])
	}
}