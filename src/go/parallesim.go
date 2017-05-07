package main

import (
	"time"
	"sync"
)

//并发：逻辑上具备同时处理多个任务的能力
//并行：物理上在同一时间执行多个任务任务

//go println("hello world")
//
//go func(s string){
//	println(s)
//}("hello,world")

//使用go关键字 即可创建并发任务
//关键字go并非执行并发操作 而是创建一个并发任务单元 会被放置在系统队列中


//var c int
//
//func counter() int{
//	c++
//	return c
//}
////goroutine会跟defer一样  延迟执行  会发生参数立即复制
//func main() {
//	a := 100
//
//	go func(x ,y int) {
//		time.Sleep(time.Second) //让goroutine在main逻辑之后执行
//		println("go:",x,y)
//	}(a,counter())   //立即计算并复制参数
//
//	a += 10
//
//	println("main:",a,counter())
//
//	time.Sleep(time.Second * 3)
//}

//func main() {
//	exit := make(chan struct{})  //创建通道 因为仅是通知 数据并没有实际意义
//
//	go func() {
//		time.Sleep(time.Second)
//		println("goroutine done.")
//
//		close(exit)   //关闭通道发出信号
//	}()
//
//
//	println("main ...")
//
//	<-exit   //如收到通道关闭的通知 就解除阻塞 否则处于等待状态
//
//	println("main exit.")
//}

//等待多个任务完成 可以采用sync.waitgroup 通过计数器 递减 直到为零 退出

//func main() {
//	var wg sync.WaitGroup
//
//	for i:=0;i<10;i++ {
//		wg.Add(1)
//
//		go func(id int) {
//			defer wg.Done()
//			time.Sleep(time.Second)
//			println("goroutine",id,"done.")
//		}(i)
//
//	}
//	println("main ...")
//
//	wg.Wait()  //阻塞 直到计数器为零
//
//
//	println("main exit.")
//}

//可在多处使用Wait阻塞 它们都能收到通知

func main() {
	var wg sync.WaitGroup

	wg.Add(1)


	go func() {
		wg.Wait()  //阻塞 等待通知
		println("wait exit.")
	}()

	go func() {
		time.Sleep(time.Second)
		println("done.")
		wg.Done()
	}()

	wg.Wait()  //阻塞 等待通知

	println("main exit")
}