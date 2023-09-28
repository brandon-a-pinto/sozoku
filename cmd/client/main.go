package main

import (
	"github.com/brandon-a-pinto/sozoku/internal/client/connection"
)

func main() {
	conn, err := connection.ConnectWithServer("127.0.0.1", "4444")
	if err != nil {
		panic(err)
	}

	defer conn.Close()
}
