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

	//鉴于通道本身 就是一个并发安全的队列 可用作ID generator Pool等用途

}

type pool chan []byte

func newPool(cap int) pool {
	return make(chan []byte,cap)
}

func (p pool) get() []byte{
	var v []byte

	select {
	case v<-p:  	//返回
	default:	//返回失败 新建
		v = make([]byte,10)
	}

	return v
}

func (p pool) put(b []byte) {
	select {
	case p<-b:		//放回
	default:		//放回失败
	}
}