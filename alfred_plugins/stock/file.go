package main

import (
	"io/ioutil"
	"os"
)


func ReadFromFile(filePath string)string{
	//f, err :=os.Open(filePath)
	//if err!=nil {
	//	fmt.Println("打开["+filePath+"]句柄失败,%v", err)
	//	return ""
	//}
	//info, _ := f.Stat()
	////如果文件大于100M则报错
	//if info.Size()> (2<<19)*100{
	//	fmt.Println("["+filePath+"]太大，内存会炸锅的,%v", err)
	//	return ""
	//}
	//fmt.Println("文件大小为",info.Size())
	//
	//defer f.Close()
	//一次性加载
	b, _ := ioutil.ReadFile(filePath) // just pass the file name
	//if err != nil {
	//	fmt.Print(err)
	//}

	str := string(b) // convert content to a 'string'
	return str

	//逐行读
	//var buffer = ""
	//scanner := bufio.NewScanner(f)
	//for scanner.Scan() {
	//	lineText := scanner.Text()
	//	buffer+=(lineText)
	//
	//}
	//return buffer


	//使用缓存的写法
	//bufReader := bufio.NewReader(f)
	////创建一个1024byte大小的缓存(如果这个缓存比文件大小还要小，那会导致文件读取不完整)
	//buf := make([]byte, info.Size()+100)
	//
	//for {
	//	readNum, err := bufReader.Read(buf)
	//	if err != nil && err != io.EOF {
	//		panic(err)
	//	}
	//	if 0 == readNum {
	//		break
	//	}
	//}
	//return string(buf)

}




func ExportToFile(filePath string, data string) {
	fp, err := os.Create(filePath) // 创建文件句柄
	if err != nil {
		//fmt.Println("创建文件["+filePath+"]句柄失败,%v", err)
		return
	}
	defer fp.Close()
	fp.WriteString(data) // 写入UTF-8 BOM
	//w := csv.NewWriter(fp) //创建一个新的写入文件流
	//w.WriteAll([][]string{[]string{data}})
	//w.Flush()
}
