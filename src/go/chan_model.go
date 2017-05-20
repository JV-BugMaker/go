package main

import "sync"
//模式

//通常使用工厂方法将goroutine和通道绑定

type reveiver struct {
	sync.WaitGroup
	data chan int
}


func newReceiver() *reveiver{
	r := &reveiver{
		data:make(chan int),
	}

	r.Add(1)
	go func() {
		defer r.Done()

		for x:=range r.data{
			println("recv:",x) //接收到消息 直到通道关闭
		}
	}()

	return r
}

func main() {
	r := newReceiver()
	r.data <- 1
	r.data <- 2

	close(r.data)		//关闭通道 发出结束通知
	r.Wait()		//等待接收者处理结束
}