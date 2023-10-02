package execution

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"net"
	"os"
)

type Command struct {
	Output string
	Error  string
}

func ExecuteCommand(connection net.Conn) {
	reader := bufio.NewReader(os.Stdin)
	loop := true

	for loop {
		fmt.Printf("-> ")
		cmd, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}

		connection.Write([]byte(cmd))
		if cmd == "stop\n" {
			loop = false
			continue
		}

		cmdStruct := &Command{}
		decoder := gob.NewDecoder(connection)
		err = decoder.Decode(cmdStruct)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(cmdStruct.Output)
		if cmdStruct.Error != "" {
			fmt.Println(cmdStruct.Error)
		}
	}
}
