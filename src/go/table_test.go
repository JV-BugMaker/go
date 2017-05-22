package main

import (
	"testing"
)
//可以使用一种类似数据表的模式来批量输入条件并依次比对结果

func add(x,y int) int{
	return x+y
}

func TestAdd(t *testing.T){
	var tests = []struct{
		x int
		y int
		expect int
	}{
		{1,1,2},
		{2,2,4},
		{3,2,5},
	}

	for _,tt := range tests{
		actual := add(tt.x,tt.y)
		if actual != tt.expect{
			t.Errorf("Add (%d,%d):expect %d,actual %d",tt.x,tt.y,tt.expect,actual)
		}
	}
}

//借助MainStart自行构建M对象 通过命令行参数相配合 实现不同的测试组合
func TestMain(m *testing.M){
	//setup
	//code := m.Run()  //调用测试用例函数

	//tearDown
	//os.Exit(code)   //注意os.Exit 不会执行 defer

	match := func(pat,str string) (bool,error){   //pat 命令行参数 -run 提供的过滤条件
		return true,nil				//str：InternalTest.Name
	}

	tests := []testing.InternalTest{
		{"b":TestB},
		{"a":TestA},
	}

	benchmarks := []testing.InternalBenchmark{}
	examples := []testing.InternalExample{}

	m = testing.MainStart(match,tests,benchmarks,examples)
	os.Exit(m.Run())


}


