package main

import "log"

func main() {
	defer func() {
		if err:= recover();err !=nil{
			log.Fatalln(err)
		}
	}()

	panic("test") //中断流程 执行延迟函数
	println("panic") //永远不会执行
}
