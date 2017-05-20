package main
//同步 通道和锁 概念

type data struct {
	sync.Mutex  //作为匿名字段时 相关方法必须实现为pointer-receiver
}

func (d data) test(s string){
	d.Lock()

	defer d.Unlock()

	for i:=0;i < 5 ;i++{
		println(s,i)
		time.Sleep(time.Second)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	var d data
	go func(){
		defer wg.Done()
		d.test("read")
	}()

	go func(){
		defer wg.Done()
		d.test("write")
	}()

	wg.Wait()
}

func dosomething(){
	m.Lock()
	url:=cache["key"]
	m.Unlock()  //锁粒度控制在最小范围内
	http.Get(url)
}

//对性能要求比较高时，应该避免使用defer Unlock
//读写并发时，用RWMutex性能会更好一些
//对单个数据读写保护，可尝试使用原子操作
//执行严格测试 尽可能打开数据竞争检查

