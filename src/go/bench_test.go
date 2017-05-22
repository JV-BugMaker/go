package main

import (
	"testing"
)

//性能测试函数以Benchmark为名称前缀 同样保存在*_test.go


func add(x,y int){
	return x + y
}

func BenchmarkAdd(b *testing.B){
	for i:=0;i< b.N;i++{
		_ = add(1,2)
	}
}

//测试工具默认不会执行性能测试 须使用bench参数 它通过逐步调整B.N值，反复执行测试函数 直到能获得准确的测量结果




//默认就以并发方式执行测试 可使用cpu参数设定多个并发限制来观察结果

//使用benchmem参数 输出内存分配信息

//-cover 参数测试代码覆盖率
//指定covermode和coverprofile参数
//set：是否执行
//count：执行参数
//atomic:执行参数，执行并发模式
