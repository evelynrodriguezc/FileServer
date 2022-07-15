package main

import (
	"encoding/gob"
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
	err = gob.NewEncoder(c).Encode(msg)
	if err != nil {
		fmt.Println(err)
	}
	c.Close()

}

func main() {
	go client()

	var input string
	fmt.Scanln(&input)
}
