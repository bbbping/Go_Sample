package main

import (
	"fmt"
)

type Parent struct {
}

func (t Parent)ShowA(){
	fmt.Println("show A")
	t.ShowB()
}

func (t Parent)ShowB(){
	fmt.Println("show B")
}

type Son struct {
	Parent
}

func (t Son)ShowB(){
	fmt.Println("show People B")
}


func main()  {
	s:=Son{}
	s.ShowA()

}
