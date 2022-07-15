package main
import (
	"strings"
	"net"
	"fmt"
	"os"
)

const (
    HOST = "localhost"
    PORT = "80"
    TYPE = "tcp"
)

func main() {

	argumentos := os.Args[1:]
	commaSeparatedMessage := joinMessage(argumentos)

	reply := sendMessageTo(HOST, PORT, commaSeparatedMessage)
	fmt.Println(reply)
}

func joinMessage(arrayMessage []string) string{
	return strings.Join(arrayMessage, ",")
}

func splitMessage(commaSeparatedMessage string)[]string{
	return strings.Split(commaSeparatedMessage, ",")
}

func sendMessageTo(host string, port string, message string) string{
	conn, err := net.Dial("tcp", host+":"+port) // intento de conneccion al host
	if err != nil {
        println("Dial failed:", err.Error())
        os.Exit(1)
    }
	_, err = conn.Write([]byte(message)) //envio el mensaje
	if err != nil {
        println("Write to server failed:", err.Error())
        os.Exit(1)
    }

	respuesta := make([]byte, 1024) //buffer para respuesta
	_, err = conn.Read(respuesta) //leo la respuesta del servidor
	if err != nil {
        println("Response from server failed:", err.Error())
        os.Exit(1)
    }

	conn.Close() //cerrar la connecion
	return string(respuesta) //respuesta de el servidor
}