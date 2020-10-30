package main

import (
	"fmt"
)
//const 这个是声明常量的，
func main() {
	const (
		n1 = iota //这个iota的作用就是索引，是常量的计数器。定义枚举的时候很有用。
		n2 
		n3  //如果没有赋值的话，会自动使用上边的值，但是不包括iota
	);
	const (

	)
	fmt.Println(n1,n2,n3);
	const (
		a1 = iota
		a2 = 100
		a3 = iota
		a4
	)
	fmt.Println(a1,a2,a3,a4)
	const (
		b1 = 1 << (10 * iota)   // "<<" 这个符号表示1的二进制表示向左移动10位，也就是由1变成了10 000 000 000
		b2 = 1 << (10 * iota)
		b3 = 1 << (10 * iota)
	)
	fmt.Println(b1,b2,b3)
	//多个iota定义在一行
	const (
		c1,c2 = iota + 1, iota + 2	 // 1,2
		c3,c4                    	 // 2,3
		c5,c6					   	 // 3,4
	)
	fmt.Println(c1,c2,c3,c4,c5,c6)
}