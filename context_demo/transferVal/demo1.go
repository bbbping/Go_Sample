package main

import (
	"context"
	"fmt"
)

func main(){
	c := context.WithValue(context.Background(), "name", "mike")
	A(c)
}

func A(cxt context.Context){
	fmt.Println("func A")
	fmt.Println("cxt",cxt)
	cxt2 := context.WithValue(cxt, "age", "29")
	fmt.Println("cxt2",cxt2)
	B(cxt2)
}

func B(cxt context.Context){
	fmt.Println("func B")
	name := cxt.Value("name")
	age := cxt.Value("age")
	fmt.Println("name:",name,"age:",age)
	fmt.Println("cxt2",cxt)
}
