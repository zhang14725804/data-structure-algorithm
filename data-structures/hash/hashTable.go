/*
	哈希表（实现简单hash表）
*/ 
const (
	Deleted = iota // 数据已被删除
	MinTableSize = 100 // 哈希表大小
	Legimate = iota // 已存在合法数据
	Empty = iota // 数据为空
)

// 哈希函数
func MySHA(str interface{},tableSize int) int {
	h:=0
	var chars []byte
	// 字符串转化为字节数组
	if s,ok:=str.(string);ok{
		chars = []byte(s)
	}
	for _,v := range chars{
		// h = (h<<17)+int(v)
		h = (h<<17|123&1235^135)+int(v)
	}
	return h % MinTableSize
}

func MySHA256(str string,tableSize int) int {
	sha := sha256.New()
	sha.Write([]byte(str))
	b:=sha.Sum(nil)

	h:=0
	for _,v := range b{
		h = (h<<17|123&1235^135)+int(v)
	}
	return h % MinTableSize
}

// 
type HashFunc func(data interface{},tableSize int)int

type HashEntry struct{
	data interface{}
	kind int // 类型
}

// 哈希表他么什么结构
type HashTable struct{
	tableSize int
	cells []*HashEntry //
	fn HashFunc
}

type HashTableGolang interface{
	Find(data interface{}) int
	Insert(data interface{})
	Empty()
	GetValue(index int) interface{}
}

func NewHashTable(size int,fn HashFunc)(*HashTable,error){
	if size < MinTableSize{
		return nil,errors.New("哈希空间太小")
	}
	if fn == nil{
		return nil,errors.New("没有哈希函数")
	}
	//
	table := new(HashTable)
	table.tableSize = size
	table.cells = make([]*HashEntry,size)
	table.fn = fn
	for i:=0;i<table.tableSize;i++{
		table.cells[i] = new(HashEntry)
		table.cells[i].data = nil
		table.cells[i].kind = Empty
	}
	return table,nil
}

// 
func (t *HashTable) Find(data interface{}) int{
	col := 0
	pos := t.fn(data,t.tableSize) // 计算哈希位置
	if t.cells[pos].kind!=Empty && t.cells[pos].data != data{
		col += 1
		pos := 2*pos-1 // 平方探测
		if pos>t.tableSize{
			pos -= t.tableSize // 越界返回
		}
	}
	return pos
}

func (t *HashTable) Insert(data interface{}){
	pos := t.Find(data) // 
	entry := t.cells[pos]
	if entry.kind != Legimate {
		entry.kind = Legimate
		entry.data = data
	}
}

// 清空
func (t *HashTable) Empty(){
	for i:=0;i<t.tableSize;i++{
		if t.cells[i] == nil{
			continue
		}
		t.cells[i].kind = Deleted
	}
}

func (t *HashTable) GetValue(index int) interface{}{
	if index>t.tableSize{
		return nil
	}
	entry := t.cells[index]
	if entry.kind == Legimate{
		return entry.data
	}
	return nil
}
func main(){
	table,_ := NewHashTable(100,MySHA)
	table.Insert("qaz1")
	table.Insert("qaz2")
	table.Insert("qaz3")
	pos:=table.Find("qaz1")
	fmt.Println(table.GetValue(pos))
	pos=table.Find("qaz2")
	fmt.Println(table.GetValue(pos))
}

func main1() {
	fmt.Println(MySHA("abcd1",100))
	fmt.Println(MySHA("abcd2",100))
	
	fmt.Println(MySHA256("abcd1",100))
	fmt.Println(MySHA256("abcd2",100))
}