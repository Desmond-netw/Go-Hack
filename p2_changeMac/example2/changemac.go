package main

import (
	"flag"
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

	// flags initilization
	iface := flag.String("iface", "", "Interface to change mac")
	newMac := flag.String("newMac", "", "New mac address")

	flag.Parse()

	if iface == nil {
		log.Fatal("provide hardware interface")
		os.Exit(1)
	}
	execCommand("sudo", []string{"ifconfig", *iface, "down"})
	execCommand("sudo", []string{"ifconfig", *iface, "hw", "ether", *newMac})
	execCommand("sudo", []string{"ifconfig", *iface, "up"})
}
