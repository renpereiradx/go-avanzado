package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
)

var site = flag.String("site", "scanme.nmap.org", "url to scan")

func main() {
	flag.Parse()
	var wg sync.WaitGroup
	defer wg.Wait()
	for i := 0; i < 65535; i++ {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			connection, err := net.Dial("tcp", fmt.Sprintf("%s:%d", *site, port))
			if err != nil {
				return
			}
			connection.Close()
			fmt.Printf("Port %d is open\n", port)
		}(i)
	}
}

/*
	Flag in terminal
	--site=scanme.webscantest.com
	--site=localhost
*/
