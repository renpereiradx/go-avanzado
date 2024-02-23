package main

import (
	"fmt"
	"net"
)

func main() {
	for i := 0; i < 100; i++ {
		// connection, err := net.Dial("tcp", fmt.Sprintf("%s:%d", "scanme.nmap.org", i))
		connection, err := net.Dial("tcp", fmt.Sprintf("%s:%d", "localhost", i))
		if err != nil {
			continue
		}
		fmt.Println(&connection)
		defer connection.Close()
		fmt.Printf("Port %d is open\n", i)
	}
}
