package main

//通过自定义类型 实现enum的自增情况
type color byte

const (
	black color = iota
	red
	blue
)

func test(c color){
	println(c)
}
func main() {
	//借助iota标识符实现一组自增常量值来实现枚举类型
	const (
		x = iota //0
		y	//1
		z	//2
		c = 100 //中断iota的自增，需要显示转换 且后续自增值按行序递增
		d
		e = iota //5 注意这边的iota是恢复自增 但是c和d的中断可以理解为不完全中断

	)
	//不同于变量在运行期间分配存储内存（非优化状态），常量通常会被编译器在预处理阶段直接展开
	//作为指令数据使用

	//数字常量不会分配存储空间 无须象变量那样通过内存寻址来取值

	//无类型常量 直接展开 但是指定类型常量编译器会进行强类型指定

	test()
}



