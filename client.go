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
	PORT = "5000"
	TYPE = "tcp"
)

func main() {

	reply := sendMessageTo(HOST, PORT)
	fmt.Println(reply)
}

func sendMessageTo(host string, port string) string {
	conn, err := net.Dial("tcp", host+":"+port) // intento de conneccion al host
	if err != nil {
		println("Dial failed:", err.Error())
		os.Exit(1)
		defer conn.Close()
	}

	fmt.Println("now connected to", HOST+":"+PORT)

	for {
		message, err := bufio.NewReader(os.Stdin).ReadString('\n')

		if err != nil {
			log.Panic(err)
			continue
		}

		_, err = conn.Write([]byte(message)) //envio el mensaje
		if err != nil {
			println("Write to server failed:", err.Error())
			os.Exit(1)
		}

		respuesta := make([]byte, 1024) //buffer para respuesta
		_, err = conn.Read(respuesta)   //leo la respuesta del servidor
		if err != nil {
			println("Response from server failed:", err.Error())
			os.Exit(1)
		}

	}

	return "nil" //respuesta de el servidor
}
