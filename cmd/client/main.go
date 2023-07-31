package main

import (
	"log"
	"net"
)

func main() {
	ip := "127.0.0.1"
	port := "4444"
	address := ip + ":" + port

	connection, err := net.Dial("tcp", address)
	if err != nil {
		log.Printf("[-] Connection refused\n")
	} else {
		log.Printf("[+] Connection established (%s)\n", connection.RemoteAddr().String())
	}
}
