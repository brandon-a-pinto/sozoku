package main

import (
	"github.com/brandon-a-pinto/sozoku/internal/server/core"
)

func main() {
	conn, err := core.ConnectWithClient("127.0.0.1", "4444")
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	core.Options(conn)
}
