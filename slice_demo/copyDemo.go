package main

import "fmt"

func main(){
	a:= []int{1,2,3,4,5}
	b:= []int{7,8,9}
	//copy只能用于切片，不能用于数组，而且第一个参数是目标，第二是源头

	copy(a,b)
	fmt.Println(a)

	// 此外，a和b的长度都必须要先分配好,不然发现会复制不进去
	c:=[]int{}
	copy(c,a)
	fmt.Println(c)


	a1:= []int{1,2,3,4,5}
	b1:= []int{7,8,9}
	//假如目标的切片长度小于源切片，只会复制头部的过去，再覆盖

	copy(b1,a1)
	fmt.Println(b1)

}
