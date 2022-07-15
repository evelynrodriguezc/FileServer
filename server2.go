package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

func server() {
	s, err := net.Listen("tcp", ":7777")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		c, err := s.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleClient(c)
	}
}

func handleClient(c net.Conn) {
	var msg string
	err := gob.NewDecoder(c).Decode(&msg)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("Mensaje: ", msg)
	}

}

func main() {
	go server()

	var input string
	fmt.Scanln(&input)
}
