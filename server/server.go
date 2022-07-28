package main

import (
	"bufio"
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
	channels (map[string][]string) = make(map[string][]string) //channel name and [who are suscribed]
)

func main() {
	listen, err := net.Listen(TYPE, HOST+":"+PORT)
	if err != nil {
		log.Panic(err)
		os.Exit(1)
	}
	fmt.Println("now listen on", HOST+PORT)

	defer listen.Close() //defer starts an action once it's already ejecuted
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

func joinMessage(arrayMessage []string) string { // send multipart (example: to suscribe to a channel = option like: sub and name: channel name) bytes
	return strings.Join(arrayMessage, " ")
}

func splitMessage(commaSeparatedMessage string) []string { //?
	return strings.Split(commaSeparatedMessage, " ")
}

func readMessage(conn net.Conn) string {
	clientMessage, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	return clientMessage
}

func handleMessages(conn *net.Conn) { //arguments?
	fmt.Println("I'm goint to take the address")
	var remoteAddress string = (*conn).RemoteAddr().String() // takes the conn IP
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

		(*conn).Write([]byte(response + "\n")) //?
		fmt.Println(response)
	}
}
