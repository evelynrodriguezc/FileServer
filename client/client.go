package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

const (
	HOST = "localhost"
	PORT = "8081"
	TYPE = "tcp4"
)

func main() {

	reply := sendMessageTo(HOST, PORT)
	fmt.Println(reply)
}

func sendMessageTo(host string, port string) string {
	conn, err := net.Dial("tcp", host+":"+port) // host conn
	if err != nil {
		println("Dial failed:", err.Error())
		os.Exit(1)
		defer conn.Close()
	}

	fmt.Println("now connected to", HOST+":"+PORT)

	for {
		message, err := bufio.NewReader(os.Stdin).ReadString('\n') //?

		if err != nil {
			log.Panic(err)
			continue
		}
		if message == "exit" {
			break
		}

		_, err = conn.Write([]byte(message)) //sends the mss
		if err != nil {
			println("Write to server failed:", err.Error())
			os.Exit(1)
		}

		respuesta := make([]byte, 1024) //buffer for response
		_, err = conn.Read(respuesta)   //reading servers response
		if err != nil {
			println("Response from server failed:", err.Error())
			os.Exit(1)
		}
	}

	return "nil" //servers response
}
