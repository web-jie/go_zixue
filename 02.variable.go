package main

import (
	"fmt"
)

var m = 100  //全局变量

func main() {

	n := 10 // ":=" 这个符号只能在函数内使用，就是省略了定义这个词, 等价于var n = 10 
	
	m := 200   //这个是局部变量
	fmt.Println(m,n) //打印出来m,n的值
}



