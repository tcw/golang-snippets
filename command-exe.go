package main

import (
	"fmt"
	"os/exec"
	"log"
	"strings"
)

func main() {
	lookForCommand("unzip")
	output := runCommandString("mkdir target")
	fmt.Println(output)
	output = runCommandString("unzip -d ./target text.zip")
	fmt.Println(output)
	output = runCommandString("cp text.zip target/")
	fmt.Println(output)
}

func lookForCommand(command string) {
	path, err := exec.LookPath(command)
	if err != nil {
		log.Fatalf("[%s] must be installed\n", command)
	}
	fmt.Printf("Found command [%s] at %s\n", command, path)
}

func runCommandString(commandString string) string {
	cmdAndArgs := strings.Split(commandString, " ")
	return runCommand(cmdAndArgs[0], cmdAndArgs[1:])
}

func runCommand(command string, commandArgs []string) string {
	out, err := exec.Command(command, commandArgs...).Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(out)
}