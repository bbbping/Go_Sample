package main

import "fmt"

func main()  {
	var a int = 123
	var i interface{} = a

	if v,ok:= i.(int);ok {
		fmt.Println("它是int类型，值：",v)
	}

	if v,ok:= i.(string);ok {
		fmt.Println("它是string类型，值：",v)
	}
}
