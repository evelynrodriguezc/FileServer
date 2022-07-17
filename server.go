package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

const (
	HOST = "localhost"
	PORT = "5000"
	TYPE = "tcp"
)

func main() {
	listen, err := net.Listen(TYPE, HOST+":"+PORT)
	if err != nil {
		log.Panic(err)
		os.Exit(1)
	}
	fmt.Println("now listen on", HOST+PORT)

	// channels := make(map[string][]string)

	// close listener
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		defer conn.Close()
		var remoteAddress string = conn.RemoteAddr().String()
		fmt.Println(remoteAddress)
		go handleMessages(conn)
	}
}

func joinMessage(arrayMessage []string) string {
	return strings.Join(arrayMessage, ",")
}

func splitMessage(commaSeparatedMessage string) []string {
	return strings.Split(commaSeparatedMessage, ",")
}

func readMessage(conn net.Conn) string {
	buffer := make([]byte, 1024)
	_, err := conn.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}
	return string(buffer)
}

func handleMessages(conn net.Conn) {
	for {
		response := ""
		message := readMessage(conn)
		sliceMessage := splitMessage(message)
		opcion := sliceMessage[0]
		switch opcion {
		case "sub":
			// channels[opcion[2]] = append(channels[opcion[2]], opcion[1])
			// response = opcion[1] + " subscribed to " + opcion[2]
			response = "sub"
		case "chans":
			response = "getchannels"
		case "upload":
			response = "upload"
		case "receive":
			response = "download"
		default:
			response = message
		}

		conn.Write([]byte(response + "\n"))
		fmt.Println(response)
	}
}
