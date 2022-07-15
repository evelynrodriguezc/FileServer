package main
import (
	"strings"
	"log"
    "net"
    "os"
)

const (
    HOST = "localhost"
    PORT = "80"
    TYPE = "tcp"
)


func main() {
    listen, err := net.Listen(TYPE, HOST+":"+PORT)

	channels := make(map[string][]string)
	
    if err != nil {
        log.Fatal(err)
        os.Exit(1)
    }
    // close listener
    defer listen.Close()
    for {
		response := ""
		conn, message := waitForActions(listen)
		opcion := splitMessage(message)
		println("user: " + opcion[1])
		switch opcion[0] {
		case "sub":
			channels[opcion[2]] = append(channels[opcion[2]], opcion[1])
			response = opcion[1]+" subscribed to "+opcion[2]
		case "chans":
			response = "getchannels"
		case "upload":
			response = "upload"
		case "receive":
			response = "download"
		}

		conn.Write([]byte(response+"\n"))
		conn.Close()
    }
}

func joinMessage(arrayMessage []string) string{
	return strings.Join(arrayMessage, ",")
}

func splitMessage(commaSeparatedMessage string)[]string{
	return strings.Split(commaSeparatedMessage, ",")
}

func waitForActions(escucha net.Listener) (net.Conn, string){
	conn, err := escucha.Accept()
	buffer := make([]byte, 1024)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	_, err = conn.Read(buffer)
	if err != nil {
        log.Fatal(err)
    }
	return conn, string(buffer)
}
