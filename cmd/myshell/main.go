package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func findCommand(cmdName string) (string, error) {
	cleanedCMDName := strings.TrimSpace(cmdName)
	return cleanedCMDName, fmt.Errorf("%s: command not found", cleanedCMDName)
}

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	//fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage

	// Wait for user input

	for {
		fmt.Fprint(os.Stdout, "$ ")
		reader := bufio.NewReader(os.Stdin)

		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		command, err := findCommand(input)

		if err != nil {
			fmt.Println(err)
		}

		log.Println(command)

	}

}
