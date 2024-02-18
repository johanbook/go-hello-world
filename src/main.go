package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const NAME string = "aus"

func parseAndExecute(cmd string) string {
	if cmd == "exit" {
		os.Exit(0)
	}

	if cmd == "help" {
		return fmt.Sprintln("There is no help")
	}

	return fmt.Sprintf("Command not found: %s\n", cmd)
}

func main() {
	fmt.Printf("Launching %s\n", NAME)

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("%s > ", NAME)
		text, _ := reader.ReadString('\n')
		rs := parseAndExecute(strings.ReplaceAll(text, "\n", ""))
		fmt.Print(rs)
	}
}
