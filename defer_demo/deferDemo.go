package main

import "fmt"

func main()  {
	defer fmt.Println("第一个defer")
	flag:=true
	if flag{
		defer fmt.Println("打印了true部分的defer")
	}else {
		// 会发现假如有条件设置了，会导致它不会进去
		defer fmt.Println("打印了false部分的defer")
	}
	flag=false //在临走之前，改变了，也没用，因为一开始就是true，将true部分的defer func放入defer栈了
	fmt.Println("结束了")
}
