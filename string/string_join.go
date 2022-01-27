package main

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
)

func main(){
	//a:= "abc"+'d' //编译不通过，因为'd'不是普通字符串,它本质上是int32，而且''
	fmt.Println(reflect.TypeOf('d'))
	fmt.Println('d')

	b:="abc"+`zzz` //可以
	fmt.Println(b)

	for _,v:=range b{
		fmt.Println(reflect.TypeOf(v),string(v))
	}

	sb:=strings.Builder{}
	sb.WriteString("abc")
	sb.WriteString("efg")
	fmt.Println(sb.String())

	bb:=bytes.Buffer{} //相对比strings.Builder，没那么高效
	bb.WriteString("abc")
	bb.WriteString("efg")
	fmt.Println(bb.String())
}
