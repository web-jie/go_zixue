package main

import (
	"fmt"
)

func main() {
	//整型  分为两个大类 按长度分为： int8 int16 int32 int64  对应无符号整型： uint8 uint16 uint32 uint64 
	//int16对应c语言的short型 int64对应long型
	//十进制
	var a int = 10
	fmt.Println("%d \n",a) //10
	fmt.Println("%b \n",a) //1010 占位符%b表示二进制
	//八进制 以0开头
	var b int = 077
	fmt.Println("%o \n",b) //77
	//十六进制   以0x开头
	var c int = 0xff
	fmt.Println("%x \n", c) //ff
	fmt.Println("%x \n", c) //FF
}
func main()  {
	//浮点型

}