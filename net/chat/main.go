package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
)

type Client chan<- string

var (
	incomingClients = make(chan Client)
	leavingClients  = make(chan Client)
	messages        = make(chan string)
)

var (
	host = flag.String("h", "localhost", "host")
	port = flag.Int("p", 3090, "port")
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
