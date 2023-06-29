package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func execPingCommand(ip string) {
	ping := exec.Command("ping", "-c", "4", ip)
	output, err := ping.Output()

	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	fmt.Println(string(output))
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		commandStrings := strings.Split(input, " ")

		if strings.Compare(commandStrings[0], "checkinfyconnect") == 0 {
			if len(commandStrings) > 1 {
				execPingCommand(commandStrings[1])
			} else {
				fmt.Println("Please provide an IP address.")
			}
		} else if strings.Compare(commandStrings[0], "exit") == 0 {
			os.Exit(0)
		} else {
			fmt.Println("Command not supported by the CLI")
		}
	}
}

