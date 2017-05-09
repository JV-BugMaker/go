package main
//goroutine任务无法设置优先级 无法获取编号 没有局部存储
import (
	"sync"
	"fmt"
	"runtime"
)
func main() {
	var wg sync.WaitGroup
	//gs 是个数组 元素是结构体
	var gs [5]struct{  //用于实现类似TLS功能
		id int
		result int
	}

	for i:=0;i<len(gs);i++{
		wg.Add(1)

		go func(id int) { //使用参数避免闭包延迟求值
			defer wg.Done()

			gs[id].id = id
			gs[id].result = (id+1)*100
		}(i)
	}

	wg.Wait()
	fmt.Printf("%+v\n",gs)
}

//gosched 暂停 释放线程去执行其他任务 当前任务被放回队列 等待下次调度时恢复执行

func main2() {
	runtime.GOMAXPROCS(1)

	exit := make(chan struct{})

	go func() {  //任务a
		defer close(exit)  //关闭通道 发出信号

		go func() { //将任务b放在此处是为了确保a先执行
			println("b")

		}()

		for i:=0;i<4;i++{
			println("a:",i)

			if i==1{  //让出当前线程 调度执行b
				runtime.Gosched()
			}

		}
	}()

	<-exit

}

func main3() {
	exit := make(chan struct{})

	go func() {
		defer close(exit)
		defer println("a") //执行

		func(){
			defer func(){
				println("b",recover() == nil) //执行 recover 返回nil
			}()

			func(){			//在多层次调用中执行goexit
				println("c")
				runtime.Goexit() //立即终止整个调用堆栈
				println("c done") //不会执行
			}()

			println("b done") //不会执行
		}()

		println("a done") //不会执行
	}()
	<-exit

	println("main exit")
}

//如果是在main.main中执行 会等待其他任务结束 然后让进程直接崩溃
func main4() {
	for i:=0;i<2;i++{
		go func(x int){
			for n:=0;n<2;n++{
				fmt.Printf("%c:%d\n",'a'+x,n)
				time.Sleep(time.Millisecond)
			}
		}(i)
	}

	runtime.Goexit()  //等待所有任务结束
	println("main exit.")
}