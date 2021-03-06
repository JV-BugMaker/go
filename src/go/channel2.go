package main

import(
	//"fmt"
)

//go鼓励使用csp通道来代替内存共享 实现并发安全
//csp：通道是显式的，双方对数据格式和具体通道都是知道的 不关心消费方身份和数量 消息未处理就会阻塞当前端
//actor 是透明的 它不在乎数据类型及通道 只要知道接收方地址即可 默认是异步方式 不关心是否被成功被处理

//func main() {
//	done := make(chan struct{})  //结束事件
//	c := make(chan string) //数据传输通道
//
//
//	go func() {
//		print("gjw")
//		s := <-c //接收消息
//		print(s)
//		close(done) //关闭通道 作为结束通知
//	}()
//
//	c <- "hi!" //发送消息
//	<-done //阻塞 直到有数据或者管道关闭
//	print("jv")
//
//}

//同步模式 必须要有配对操作的goroutine出现 否则会一直阻塞 而异步模式在缓冲区未满或数据未读完前 不会阻塞

//func main() {
//	c := make(chan int,3)  //创建带3个缓冲槽的异步通道
//
//
//	c<-1 			//缓冲区未满 不会阻塞
//	c<-2
//
//	print(<-c)		//缓冲区尚有数据 不会阻塞
//	print(<-c)
//
//}

//func main() {
//	var a,b chan int = make(chan int,3),make(chan int)
//	var c chan bool
//
//	println(a == b)
//	println(c == nil)
//
//	fmt.Printf("%p,%d\n",a,unsafe.Sizeof(a))
//
//}
//缓冲区大小仅是内部属性 不属于类型组成部分 另外通道变量本身就是指针 可用相等操作符判断是否为同一对象

func main() {
	a,b :=make(chan int),make(chan int,3)


	b<-1
	b<-2

	println("a:",len(a),cap(a))
	println("b:",len(b),cap(b))
}
//内置函数 cap和len返回缓冲区大小和当前已缓冲数量 而对于同步通道都返回0