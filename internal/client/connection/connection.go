package connection

import (
	"fmt"
	"net"
)

func ConnectWithServer(IP, Port string) (conn net.Conn, err error) {
	address := IP + ":" + Port

	connection, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Printf("[-] Connection refused\n")
	} else {
		fmt.Printf("[+] Connection established: %s\n", connection.RemoteAddr().String())
	}

	return connection, nil
}
