package main

import "sync"
//通道默认是双向的 并不区分发送和接收端 我们可以限制收发操作的方向来获取更严谨的逻辑

//通常使用类型转换来获取单向通道

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	c := make(chan int)
	var send chan<- int = c
	var recv <-chan int = c
	//接收方
	go func() {
		defer wg.Done()

		for x := range recv {
			println(x)
		}
	}()
	//发送方
	go func() {
		defer wg.Done()
		defer close(c)

		for i := 0;i<3;i++{
			send <- i
		}
	}()

	wg.Wait()
	//不能再单向通道上做逆向操作
	//同样 close不能用于接收端 只能作用于发送端
}