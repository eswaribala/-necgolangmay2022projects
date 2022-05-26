package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	listen, err := net.Listen("tcp",
		"localhost:4000")
	if err != nil {
		log.Println("Net TCP Server not running")
	} else {
		conn, error := listen.Accept()
		if error == nil {
			log.Println("Connection accepted")
			go handleConn(conn)
		}
	}

}

func handleConn(conn net.Conn) {
	fmt.Println(time.Now())
}
