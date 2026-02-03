package main

import (
	"flag"
	"fmt"
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

	if *iface == "" {
		log.Fatal("provide hardware interface")
		os.Exit(1)
	}
	if *newMac == "" {
		log.Fatal("provide hardware interface")
		return
	}
	// Bring interface down
	if err := execCommand("sudo", []string{"ifconfig", *iface, "down"}); err != nil {
		fmt.Println(err)
	}
	// Change Mac Address
	if err := execCommand("sudo", []string{"ifconfig", *iface, "hw", "ether", *newMac}); err != nil {
		fmt.Println(err)
	}

	// Bring interface up
	if err := execCommand("sudo", []string{"ifconfig", *iface, "up"}); err != nil {
		fmt.Println(err)
	}
}
