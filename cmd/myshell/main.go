package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

var (
	funcMap        = make(map[string]GenericFunc)
	builtInFuncMap = make(map[string]GenericFunc)
	InfoLogger     *log.Logger
	WarningLogger  *log.Logger
	ErrorLogger    *log.Logger
	mutex          = &sync.Mutex{} // To safely append to funcMap concurrently
)

// Function type that accepts an arbitrary number of interface{} arguments
type GenericFunc func(...interface{})

func exit(codeStr string) {
	val, err := strconv.Atoi(codeStr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	os.Exit(val)
}

func echo(strArgs ...string) {
	fmt.Println(strings.Join(strArgs, " "))
}

func init() {
	// set up logging
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	// set up built-in commands
	builtInFuncMap["exit"] = func(args ...interface{}) {
		if len(args) > 0 {
			if codeStr, ok := args[0].(string); ok {
				exit(codeStr)
			} else {
				fmt.Println("Invalid argument type for exit")
			}
		}
	}
	builtInFuncMap["echo"] = func(args ...interface{}) {
		strArgs := make([]string, len(args))

		for i, v := range args {
			strArgs[i] = v.(string)
		}
		echo(strArgs...)
	}
}

func parseCMD(cmd string) (string, []string) {
	inputs := strings.Fields(cmd)
	command := strings.TrimSpace(inputs[0])
	var arguments []string = inputs[1:]
	InfoLogger.Printf("Command: %s Argument: %v\n", command, arguments)
	return command, arguments
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

		command, arguments := parseCMD(input)
		InfoLogger.Printf("Trying %s with %v as arguments", command, arguments)
		if builtInFunc, exists := builtInFuncMap[command]; exists {
			args := make([]interface{}, len(arguments))
			for i, v := range arguments {
				args[i] = v
			}
			builtInFunc(args...)
		} else {
			fmt.Printf("%s: command not found\n", command)
			ErrorLogger.Printf("%s: command not found\n", command)
		}
	}
}
