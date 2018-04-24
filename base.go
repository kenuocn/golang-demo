package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	value()
	variableShorter()
}

//函数
func value() {
	var a string = "张高元" //定义一个 string类型的变量.

	// 定义多个类型的变量用一个var声明.
	var (
		bb byte = '8'
		cc int    = 10
		dd bool   = true
	)

	fmt.Println(a, bb, cc, dd)
}

func variableShorter() {
	a, b, c, d, s := 3, "张高元", true, "dfr", 34

	b = "zgy" // 变量b 重新赋值.

	fmt.Println(a, b, c, d, s)
}
