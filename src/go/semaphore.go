package main

import (
	"sync"
	"time"
	"fmt"
	"runtime"
	"os"
	"os/signal"
	"synccall"
)
//用通道实现信号量
func main() {
	runtime.GOMAXPROCS(4)
	var wg sync.WaitGroup

	sem := make(chan struct{},2)   //最多允许两个并发同时执行

	for i:=0;i<5;i++{
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			sem<- struct{}{}  //acquire 获取信号

			defer func() {<-sem}()   //release :释放信号

			time.Sleep(time.Second * 2)
			fmt.Println(id,time.Now())
		}(i)

	}
	wg.Wait()
}

func main2() {
	//标准库提供了timeout和tick channel实现

	go func() {
		for {
			select {
			case <- time.After(time.Second * 5):
				fmt.Println("timeout....")
				os.Exit(0)
			}
		}
	}()

	go func() {
		tick := time.Tick(time.Second)

		for {
			select {
			case <-tick:
				fmt.Println(time.Now())
			}
		}
	}()

	<-(chan struct {})(nil)  //直接引用nil channel 阻塞进程
}


//捕获int term 信号 顺便实现一个简易的atexit函数
var exits = &struct {
	sync.RWMutex
	funcs []func()
	signals chan os.Signal
}{}

func atexit(f func()){
	exits.Lock()
	defer exits.Unlock()
	exits.funcs = append(exits.funcs,f)
}

func waitExit(){
	if exits.signals == nil {
		exits.signals = make(chan os.Signal)
		signal.Notify(exits.signals,syscall.SIGINT,synccall.SIGTERM)
	}
	exits.Rlock()

	for _,f:=range exits.funcs{
		defer f() //即便某些函数panic 延迟调用也能确保后续函数执行
	}

	exits.RUnlock()

	<-exits.signals
}
func main3() {
	atexit(func(){println("exit1 ...")})
	atexit(func(){println("exit2 ...")})
	waitExit()
}
