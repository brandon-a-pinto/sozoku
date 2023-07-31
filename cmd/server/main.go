package main

import (
	"flag"
	"log"
	"net"
)

func main() {
	var connection net.Conn

	ip := flag.String("i", "127.0.0.1", "IP")
	port := flag.String("p", "4444", "Port")
	flag.Parse()

	address := *ip + ":" + *port

	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Server running (%s:%s)\n", *ip, *port)

	connection, err = listener.Accept()
	if err != nil {
		log.Printf("[-] Connection refused\n")
	} else {
		log.Printf("[+] Connection established (%s)\n", connection.RemoteAddr().String())
	}
}
