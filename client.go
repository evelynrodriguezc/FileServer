package main

import (
	"fmt"
	"net"
)

func client() {
	c, err := net.Dial("tcp", ":7777")
	if err != nil {
		fmt.Println(err)
		return
	}
	msg := "Hola mundo"
	fmt.Println(msg)
	c.Write([]byte(msg))
	c.Close()

}

func main() {
	go client()

	var input string
	fmt.Scanln(&input)
}
