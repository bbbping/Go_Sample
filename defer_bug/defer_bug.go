package main

import "fmt"

func Bug() (i int) {
	defer func() {
			i += 1
		}()

	i = 11
	return
}

func NotBug() int {

	i := 11
	defer func() {
		i += 1
	}()
	return i
}

func main() {
	fmt.Println(Bug())
	fmt.Println(NotBug())
}
