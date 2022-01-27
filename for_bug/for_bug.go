package main

import (
	"fmt"
	"time"
)

// 其实就是在使用闭包的时候，会存在线程不安全的问题
func ForBug() {
	list := []int{1, 2, 3, 4, 5}
	for _, v := range list {
		go func() {
			fmt.Println("id:", v)
		}()
	}
}

func ForNotBug() {
	list := []int{1, 2, 3, 4, 5}
	for _, v := range list {
		go func(id int) {
			fmt.Println("id:", id)
		}(v)
	}
}

// 其实就是在使用闭包的时候，会存在线程不安全的问题
func ForBug2() {
	list := []int{1, 2, 3, 4, 5}
	for i, _ := range list {
		go func() {
			fmt.Println("id:", list[i])
		}()
	}
}

func ForNotBug2() {
	list := []int{1, 2, 3, 4, 5}
	for i, _ := range list {
		go func(index int) {
			fmt.Println("id:", list[index])
		}(i)
	}
}



//实际上是因为结构体假如本身不是指针，但是使用goroutine调用了指针方法，
// 会因为v是一个中间的临时变量,使用的内存地址都相同，因此在并发环境下调用指针方法，是会出现问题的，而goroutine因为启动慢，发现最后都指向相同的位置了
func ForNotBug3() {
	list := []*PowStruct{&PowStruct{1}, &PowStruct{2}, &PowStruct{3}, &PowStruct{4}, &PowStruct{5},}
	for _, v := range list {
		fmt.Println(fmt.Sprintf("%p",&v))
		go v.Calculate()
	}
}

func ForBug3() {
	list := []PowStruct{PowStruct{1}, PowStruct{2}, PowStruct{3}, PowStruct{4}, PowStruct{5},}
	for _, v := range list {
		fmt.Println(fmt.Sprintf("%p",&v))
		go v.Calculate()
	}
}
type PowStruct struct {
	param int
}

func (p *PowStruct) Calculate() {
	fmt.Println(p.param, "*", p.param, "=", p.param*p.param)
}

func (p PowStruct) Calculate2() {
	fmt.Println(p.param, "*", p.param, "=", p.param*p.param)
}



func main() {
	//ForBug()
	//ForNotBug()
	//ForBug2()
	//ForNotBug2()
	ForNotBug3()
	//ForBug3()
	time.Sleep(1 * time.Second)
}
