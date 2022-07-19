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
	PORT = "8081"
	TYPE = "tcp4"
)

var (
	channels (map[string][]string) = make(map[string][]string)
)

func main() {
	listen, err := net.Listen(TYPE, HOST+":"+PORT)
	if err != nil {
		log.Panic(err)
		os.Exit(1)
	}
	fmt.Println("now listen on", HOST+PORT)

	defer listen.Close()
	for {
		fmt.Println("ready to listen")
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		defer conn.Close()

		go handleMessages(&conn)
	}
}

func joinMessage(arrayMessage []string) string {
	return strings.Join(arrayMessage, ",")
}

func splitMessage(commaSeparatedMessage string) []string {
	return strings.Split(commaSeparatedMessage, " ")
}

func readMessage(conn net.Conn) string {
	buffer := make([]byte, 1024)
	_, err := conn.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}
	return string(buffer)
}

func handleMessages(conn *net.Conn) {
	fmt.Println("I'm goint to take the address")
	var remoteAddress string = (*conn).RemoteAddr().String()
	fmt.Println(remoteAddress)
	for {
		response := ""
		message := readMessage(*conn)
		sliceMessage := splitMessage(message)
		opcion := strings.TrimSpace(sliceMessage[0])
		fmt.Println("opcion: ", opcion)
		if opcion == "sub" {
			channelName := sliceMessage[1]
			fmt.Printf("I'm going to suscribe to the channel %s \n", channelName)
			channels[channelName] = append(channels[channelName], remoteAddress)
			response = "sub"
		} else if opcion == "paths" {
			fmt.Println("I'm going to show the channels")
			for chann := range channels {
				fmt.Printf("%s: ", chann)
				for _, address := range channels[chann] {
					fmt.Printf(" %s", address)
				}
				fmt.Printf("\n")
			}
			response = "getchannels"
		} else {
			response = message
		}

		(*conn).Write([]byte(response + "\n"))
		fmt.Println(response)
	}
}
