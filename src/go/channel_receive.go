package main

func main() {

	done := make(chan struct{})
	c := make(chan int)

	go func() {
		defer close(done)

		//使用ok-idom方式循环接收数据方式
		for {
			x,ok := <-c
			if !ok {
				return
			}

			println(x)
		}
		//使用range方式更加简洁
		//使用ok-idom方式 能够避免很多错误问题
		for x:=range c{
			println(x)
		}

	}()

	c<-1
	c<-2
	c<-3

	//使用close及时关闭通道 引发结束通知 否则可能会导致死锁
	close(c)
	<-done
}

func main2() {
	var wg sync.WaitGroup

	ready := make(chan struct{})

	for i:=0;i<3;i++{
		wg.Add(1)

		go func() {
			defer wg.Done()

			println(id,":ready.")  //运动员 准备就绪

			<-ready  //等待发令

			println(id,":running...")
		}(i)

	}

	time.Sleep(time.Second)
	println("Ready?Go!")

	close(ready)

	wg.Wait()
	//0:ready 2:ready 1:ready
	//ready go
	//1:running
	//2:running
	//3:running
}
//还可以使用sync.Cond实现单广播或广播事件
//对于closed或nil通道 发送和接收操作都有相应规则:
//* 向已关闭通道发送数据,引发panic。
//* 从已关闭接收数据，返回已缓冲数据或零值
//* 无论收发,nil通道都会阻塞

func main3() {
	c := make(chan int,3)

	c<-10
	c<-20

	close(c)

	for i:=0;i<cpa(c)+1;i++{
		x , ok := <-c
		println(i,":",ok,x)
	}
}


//重复关闭 或关闭nil通道都会引发panic错误



