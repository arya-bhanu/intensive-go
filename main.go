package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	fmt.Println("Hello World")
	list, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen port 9000 %v\n", err)
	}
	InitGRPC(list)
}

func InitGRPC(listen net.Listener) {

}
