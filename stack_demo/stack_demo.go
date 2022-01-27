package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	Foo()
}

func Foo() {
	fmt.Printf("我是 %s方法, 谁在调用我?\n是：%s\n", printMyName(), printCallerName())
	Bar()
}

func Bar() {
	fmt.Printf("我是 %s方法, 谁又在调用我?\n是：%s\n", printMyName(), printCallerName())
	//trace()
	trace2()
	//DumpStacks()
	//GoID()
}

//打印当前方法的名称
func printMyName() string {
	//Caller可以返回函数调用栈的某一层的程序计数器、文件信息、行号。
	//0 代表当前函数，也是调用runtime.Caller的函数。1 代表上一层调用者，以此类推。
	//因为我这个方法printMyName()套了一层，所以我想打印的是调用printMyName()的方法的名称，因此需要+1
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}

//打印我这个方法被谁调用了
func printCallerName() string {
	//因为我这个方法printCallerName()套了一层，所以我想打印的是调用printCallerName()的方法的再上一层的名称，因此需要+2
	pc, _, _, _ := runtime.Caller(2)
	//FuncForPC 是一个有趣的函数， 它可以把程序计数器地址对应的函数的信息获取出来。如果因为内联程序计数器对应多个函数，它返回最外面的函数。
	//它的返回值是一个*Func类型的值，通过*Func可以获得函数地址、文件行、函数名等信息。
	return runtime.FuncForPC(pc).Name()
}

// 打印详细的堆栈，假如这个方法被内联了，会显示上一个方法的名称
func trace() {
	pc := make([]uintptr, 10) // at least 1 entry needed
	n := runtime.Callers(0, pc)
	for i := 0; i < n; i++ {
		f := runtime.FuncForPC(pc[i])
		file, line := f.FileLine(pc[i])
		fmt.Printf("%s:%d %s\n", file, line, f.Name())
	}
}

func trace2() {
	pc := make([]uintptr, 10) // at least 1 entry needed
	n := runtime.Callers(0, pc)
	//CallersFrames函数，省去遍历调用FuncForPC来获取整个堆栈
	frames := runtime.CallersFrames(pc[:n])
	for {
		frame, more := frames.Next()
		fmt.Printf("%s:%d %s\n", frame.File, frame.Line, frame.Function)
		if !more {
			break
		}
	}
}

// 打印详细的堆栈，假如这个方法被内联了，也不会受影响
func DumpStacks() {
	buf := make([]byte, 16384)
	buf = buf[:runtime.Stack(buf, true)]
	fmt.Printf("=== BEGIN goroutine stack dump ===\n%s\n=== END goroutine stack dump ===", buf)
}

//利用堆栈信息还可以获取goroutine的id,
func GoID() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}
