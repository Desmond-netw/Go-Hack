package main

import (
	"log"
	"os"
	"os/exec"
)

// system command func
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

/* --------- */

func main() {
	// ifconfig eth0 down
	// ifconfig eth0 hw ether 00:11:22:33:44:55
	//ifconfig eth0 up

	execCommand("sudo", []string{"ifconfig", "eth0", "down"})
	execCommand("sudo", []string{"ifconfig", "eth0", "hw", "ether", "00:11:22:33:44:55"})
	execCommand("sudo", []string{"ifconfig", "eth0", "up"})
}
