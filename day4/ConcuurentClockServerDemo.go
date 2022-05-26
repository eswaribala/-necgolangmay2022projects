package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listen, err := net.Listen("tcp",
		"localhost:8000")
	for {
		if err != nil {
			log.Println("Net TCP Server not running")
			continue
		} else {
			conn, error := listen.Accept()
			//fmt.Println("Server Ready.....@ tcp 4000")
			if error == nil {
				log.Println("Connection accepted")
				go handleConn(conn)
			}
		}
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	for {
		_, err := io.WriteString(conn, time.Now().Format("14:20:20"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}
