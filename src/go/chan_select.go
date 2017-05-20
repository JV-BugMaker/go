package main

import (
	"sync"
)

func main() {
	//如果同时要处理多个通道 可选用select语句 它会随机选择一个可用通道做收发操作

	var wg sync.WaitGroup
	wg.Add(2)

	a,b := make(chan int),make(chan int)

	go func() {
		defer wg.Done()


		for {
			var (
				name string
				x int
				ok bool
			)

			select {			//随机选择可用channel接收数据
			case x,ok= <-a:
				name = "a"
			case x,ok=<-b:
				name = "b"
			default:
				//处理一些默认逻辑 比如说当通道满了 生成新的通道来缓存通道
			}

			if !ok {				//如果任一通道关闭 则终止接收
				return
			}

			println(name,x)  //输出接收数据
		}

	}()

	go func() {
		defer wg.Done()
		defer close(a)
		defer close(b)

		for i:=0;i<10;i++{
			select {	//随机选择发送
			case a<-i:
			case b<-i * 10:
			}
		}
	}()

	wg.Wait()
	//即使是同一个通道 也会随机选择case执行
	//当所有通道都不可用时 select会执行default语句 如此可避开select阻塞 但是必须处理外层循环 以免陷入空耗

}
