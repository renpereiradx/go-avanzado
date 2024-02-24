package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
)

type Client chan<- string

var (
	incomingClients = make(chan Client)
	leavingClients  = make(chan Client)
	messages        = make(chan string)
)

var (
	host = flag.String("host", "localhost", "host")
	port = flag.Int("port", 3090, "port")
)

func HandleConnection(connection net.Conn) {
	defer connection.Close()
	message := make(chan string)
	go WriteMessage(connection, message)
	clientName := connection.RemoteAddr().String()

	message <- fmt.Sprintf("Welcome to the server, your name %s\n", clientName)
	messages <- fmt.Sprintf("New client is here, name %s\n", clientName)
	incomingClients <- message

	inputMessage := bufio.NewScanner(connection)
	for inputMessage.Scan() {
		messages <- fmt.Sprintf("%s: %s\n", clientName, inputMessage.Text())
	}

	leavingClients <- message
	messages <- fmt.Sprintf("%s said goodbye!", clientName)
}

func WriteMessage(connection net.Conn, messages <-chan string) {
	for message := range messages {
		fmt.Fprintln(connection, message)
	}
}

func Broadcast() {
	clients := make((map[Client]bool))
	for {
		select {
		case message := <-messages:
			for client := range clients {
				client <- message
			}
		case newClient := <-incomingClients:
			clients[newClient] = true
		case leavingClient := <-leavingClients:
			delete(clients, leavingClient)
			close(leavingClient)
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *host, *port))
	if err != nil {
		log.Fatal(err)
	}
	go Broadcast()
	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go HandleConnection(connection)
	}
}
