package connection

import (
	"fmt"
	"net"
)

func ConnectWithClient(IP, Port string) (conn net.Conn, err error) {
	address := IP + ":" + Port

	listener, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}

	fmt.Printf("[+] Server running: %s:%s\n", IP, Port)

	connection, err := listener.Accept()
	if err != nil {
		fmt.Printf("[-] Connection refused\n")
	} else {
		fmt.Printf("[+] Connection established: %s\n", connection.RemoteAddr().String())
	}

	return connection, nil
}
