package main

//测试标准库 testing
//测试代码必须放在当前包以"_test.go"结尾的文件中
//测试函数以Test为名称前缀
//测试命令(go test)忽略以"_"或"."开头的测试文件
//正常编译操作会忽略测试文件

import (
	"testing"
	"time"
	"os"
)

func add(x,y int) int{
	return x+y
}

func TestAdd(t *testing.T){
	if add(1,2)!=3{
		t.FailNow()
	}
}

//testing提供了专用类型T来控制测试结果和行为
/*
* Fail  失败:继续执行当前测试
* FailNow  失败：立即终止执行当前测试函数  Failed
* SkipNow  跳过: 停止执行当前测试函数
* Log   输出错误信息.仅失败或-v时输出 Logf
* Parallel  与同样设置的测试函数并行执行
* Error Fail+Log  Errorf
* Fatal  FailNow + Log  Fatalf
 */

//使用Parallel 可有效利用多核并行优势 缩短测试时间

func TestA(t *testing.T){
	t.Parallel()
	time.Sleep(time.Second * 2)
}

func TestB(t *testing.T){
	if os.Args[len(os.Args)-1] == "b"{
		t.Parallel()
	}
	time.Sleep(time.Second * 2)
}
//go test 执行参数 parallel 参数必须大于1

/*
* -args 命令行参数
* -v 输出详细信息
* -parallel 并发执行 默认为GOMAXPROCS -parallel 2
* -run 指定测试函数 正则表达式 -run "Add"
* -timeout 全部测试累计时间超时引发panic
* -count 重复测试次数 默认值为1
 */

