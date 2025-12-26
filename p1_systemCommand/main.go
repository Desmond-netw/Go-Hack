package main

import (
	"log"
	"os"
	"os/exec"
)

// refactor code to
func execCommand(command string, argArr []string) (err error) {
	// init arg
	arg := argArr

	// use command obj
	obj_cmd := exec.Command(command, arg...)
	obj_cmd.Stdout = os.Stdout
	obj_cmd.Stderr = os.Stderr

	// run command
	err = obj_cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	return nil

}

func main() {
	// command
	command := "ls"
	arg := []string{"-ls"}

	// call execommand
	execCommand(command, arg)

}
