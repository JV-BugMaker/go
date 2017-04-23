package main

import "sync"

//匿名字段
type data struct {
	sync.Mutex //匿名字段 由编译器查找
	buf [1024]byte
}

func main() {
	d := data{}
	d.Lock()  //编译器会处理为 sync.(*Mutex).Lock()调用
	defer d.Unlock()
}