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


	close(c)
	<-done
}
