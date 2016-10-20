package main

import (
	"fmt"
	"os/exec"
	"log"
	"os"
)

func main() {
	command := os.Args[1]
	lookForCommand(command)
	runCommand("ls","-h","-l")
}

func lookForCommand(command string) {
	path, err := exec.LookPath(command)
	if err != nil {
		log.Fatalf("[%s] must be installed\n", command)
	}
	fmt.Printf("Found command [%s] at %s\n", command, path)
}

func runCommand(command string,commandArgs ...string){
	out, err := exec.Command(command,commandArgs...).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(out))
}