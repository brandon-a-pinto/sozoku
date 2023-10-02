package main

import (
	"github.com/brandon-a-pinto/sozoku/internal/client/core"
)

func main() {
	conn, err := core.ConnectWithServer("127.0.0.1", "4444")
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	core.Options(conn)
}
