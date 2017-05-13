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
		
	}()

	c<-1
	c<-2
	c<-3


	close(c)
	<-done
}
