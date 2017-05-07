package main

import "time"
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

func main() {
	exit := make(chan struct{})  //创建通道 因为仅是通知 数据并没有实际意义

	go func() {
		time.Sleep(time.Second)
		println("goroutine done.")

		close(exit)   //关闭通道发出信号
	}()


	println("main ...")

	<-exit   //如收到通道关闭的通知 就解除阻塞 否则处于等待状态

	println("main exit.")
}