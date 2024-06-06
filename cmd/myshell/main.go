package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

func init() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func findCommand(cmdName string) (string, error) {
	cleanedCMDName := strings.TrimSpace(cmdName)
	fmt.Printf("%s: command not found\n", cleanedCMDName)
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
			ErrorLogger.Println(err)
		}

		command, err := findCommand(input)

		if err != nil {
			ErrorLogger.Println(err)
		}

		InfoLogger.Println(command)

	}

}
