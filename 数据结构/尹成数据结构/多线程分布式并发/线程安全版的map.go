type SyncMap struct{
	mymap map[string]string
	*sync.Mutex // 线程锁
}

var smap SyncMap
var done chan bool // 是否完成

func write1()  {
	keys:=[]string{"1","2","3"}
	for _,k:=range keys{
		smap.Lock()
		smap.mymap[k] = k
		smap.Unlock()
		time.Sleep(1*time.Second)
	}
	done <- true
}
func write2()  {
	keys:=[]string{"11","22","33"}
	for _,k:=range keys{
		smap.Lock()
		smap.mymap[k] = k
		smap.Unlock()
		time.Sleep(1*time.Second)
	}
	done <- true
}

func read()  {
	smap.Lock()
	fmt.Println("read-sync")
	for k,v:=range smap.mymap{
		fmt.Println(k,v)
	}
	smap.Unlock()
}

func main()  {
	smap = SyncMap{make(map[string]string),new(sync.Mutex)}
	done = make(chan bool,2)
	go write1()
	go write2()

	for{
		read()
		if len(done) == 2{
			for k,v:=range smap.mymap{
				fmt.Println(k,v)
			}
			break
		}else{
			time.Sleep(1*time.Second)
		}
	}
}