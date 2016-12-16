package main

import "flag"

import "os"
import "net"

import "log"
import "bufio"
import "fmt"

func main() {
	var isHost bool

	flag.BoolVar(&isHost, "listen", false, "Listens on the specified IP address")
	flag.Parse()

	if isHost {
		// go run main.go -listen <ip>
		connIP := os.Args[2]
		runHost(connIP)
	} else {
		// go run main.go <ip>
		connIP := os.Args[1]
		runGuest(connIP)
	}
}

const port = "8080"

func runHost(ip string) {
	ipAndPort := ip + ":" + port
	listener, listenErr := net.Listen("tcp", ipAndPort)

	if listenErr != nil {
		log.Fatal("Error: ", listenErr)
	}

	conn, acceptErr := listener.Accept()

	if acceptErr != nil {
		log.Fatal("Error: ", acceptErr)
	}

	reader := bufio.NewReader(conn)
	message, readErr := reader.ReadString('\n')

	if readErr != nil {
		log.Fatal("Error: ", readErr)
	}

	fmt.Println("Message received: ", message)
}

func runGuest(ip string) {

}
