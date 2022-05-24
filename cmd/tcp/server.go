package main

import (
	"log"
	"net"
	"time"
)

type User struct {
	ID             int
	Addr           string
	EnterAt        time.Time
	MessageChannel chan string
}

func main() {
	// 以下是socket编程的框架
	listener, err := net.Listen("tcp", ":2020") //net 包用于进行监听的结果
	if err != nil {
		panic(err) //查看监听有没有错误
	}

	go broadcaster()

	for {
		conn, err := listener.Accept() //conn 是从端口接受到的信息
		if err != nil {
			log.Println(err)
			continue
		}

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	// 1.建立新用户，构建属于他的实例
	user := &User{}
}
