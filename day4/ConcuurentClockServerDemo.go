package main

import (
	"io"
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
	defer conn.Close()
	for {
		_, err := io.WriteString(conn, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}
