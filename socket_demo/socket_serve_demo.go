package main

import (
	"log"
	"net"
)

//一个简单的socket编程例子
func main() {
	// Part 1: create a listener，创建一个tcp监听
	l, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("Error listener returned: %s", err)
	}
	defer l.Close()

	for {
		// Part 2: accept new connection 然后轮询这个监听
		c, err := l.Accept()
		if err != nil {
			log.Fatalf("Error to accept new connection: %s", err)
		}

		// Part 3: create a goroutine that reads and write back data
		// 创建一条协程读取这个tcp端口的数据
		go func() {
			log.Printf("TCP session open")
			defer c.Close()

			for {
				d := make([]byte, 100)

				// Read from TCP buffer
				_, err := c.Read(d)
				if err != nil {
					log.Printf("Error reading TCP session: %s", err)
					break
				}
				log.Printf("reading data from client: %s\n", string(d))

				// write back data to TCP client
				_, err = c.Write(d)
				if err != nil {
					log.Printf("Error writing TCP session: %s", err)
					break
				}
			}
		}()
	}
}