package main

import (
	"fmt"
	"time"
	"sync"
)

//字典

//作为无序集合 字典要求key必须支持相等运算符(== !=)的数据结构，比如说,数字、字符串、指针、数组、结构体以及对应接口类型

func main() {
	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2


	m2 := map[int]struct{
		x int

	}{
		1:{x:100},
		2:{x:200},
	}

	fmt.Println(m,m2)
	loop()
}

func op(){
	//string是key的类型 int是value的类型
	m := map[string]int{
		"a":1,
		"b":2,
	}

	m["a"] = 10 //修改
	m["c"] = 30 //新增

	if v,ok := m["d"];ok{
		print(v)
	}

	delete(m,"d")  //删除键值对 不存在的时候 会报错
}
//map 迭代

func loop(){
	m := make(map[string]int)

	for i := 0;i < 8;i++{
		m[string('a'+i)] = i
	}

	fmt.Println(m)

	for i := 0;i < 4;i++ {
		for k,v := range m {
			print(k,":",v," ")
		}
	}

	println()
}

//len返回当前map的key数量 cap不接受map类型
//字典是not addressable 故不能直接修改value成员(结构或者数组)

func error(){
	type user struct {
		name string
		age int
	}

	m := map[int]user{
		1:{"Tom",19},
	}

	m[1].age += 1 //错误
}
//正确做法是返回整个value，待修改后再设置字典键值 或者直接用指针类型
type user struct {
	name string
	age int
}
func success(){
	 m := map[int]user{
		 1:{"Tome",19},
	 }
	//通过复制变量
	u := m[1]

	u.age += 1
	m[1] = u
	//使用指针
	m2 := map[int]*user{
		1:&user{"Jack",19},
	}

	m2[1].age += 1
}

//不能对nil字典进行写操作 但却能读


//安全 在迭代期间删除或新增键值是安全的

func deleteKey(){
	m := make(map[int]int)

	for i:=0;i<10;i++{
		m[i] = i
	}

	for k := range m {
		if k == 5 {
			m[100] = 100
		}
		delete(m,k)
		fmt.Println(k,m)
	}

}
//运行时会对map进行并发操作检测 。也就是说进程对map操作会锁住这个map

func check(){
	m := make(map[string]int)

	//匿名函数
	go func(){
		for {
			m["a"] += 1   //写操作
			time.Sleep(time.Microsecond)
		}
	}()
	//并发操作
	go func(){
		for{
			_=m["b"]   //读操作
			time.Sleep(time.Microsecond)
		}
	}()

	select {

	} //阻止进程退出
}
//可以启用 数据竞争 检查此类问题 go run -race test.go
//可以使用sync.RWMutex实现同步，避免读写操作同时进行

func syncFunc(){
	var lock sync.RWMutex  //使用读写锁，以获取最佳性能
	m := make(map[string]int)

	go func(){
		for {
			lock.Lock() //注意锁的粒度
			m["a"] += 1
			lock.Unlock() //不能使用defer

			time.Sleep(time.Microsecond)
		}
	}()

	go func() {
		for {
			lock.RLock() //读锁
			_ = m["b"]
			lock.RUnlock()

			time.Sleep(time.Microsecond)
		}
	}()

	select {

	}
}

//字典对象本身就是指针包装 传参数时无需再次获取地址

//对于数据量小的 建议直接使用字典存储拷贝 而非指针
//字典不会收缩内存 所以适当替换成新对象是必要的
