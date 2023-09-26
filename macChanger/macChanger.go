package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/exec"
)

func main() {
	interfaceName := flag.String("i", "", "Network interface name (e.g., eth0)")
	newMAC := flag.String("m", "", "New MAC address (e.g., 00:11:22:33:44:55)")
	flag.Parse()

	if *interfaceName == "" || *newMAC == "" {
		fmt.Println("Usage: macChanger -i [Network Interface Name] -m [New MAC Address]")
		os.Exit(1)
	}

	interfaceObj, err := net.InterfaceByName(*interfaceName)
	if err != nil {
		fmt.Printf("Network interface not found: %v\n", err)
		os.Exit(1)
	}

	err = setMACAddress(interfaceObj, *newMAC)
	if err != nil {
		fmt.Printf("Failed to change MAC address: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("MAC address successfully changed to: %s\n", *newMAC)
}

func setMACAddress(iface *net.Interface, newMAC string) error {
	err := disableInterface(iface)
	if err != nil {
		return err
	}

	err = setNewMACAddress(iface, newMAC)
	if err != nil {
		return err
	}

	err = enableInterface(iface)
	if err != nil {
		return err
	}

	return nil
}

func disableInterface(iface *net.Interface) error {
	err := exec.Command("ifconfig", iface.Name, "down").Run()
	if err != nil {
		return err
	}
	return nil
}

func setNewMACAddress(iface *net.Interface, newMAC string) error {
	err := exec.Command("ifconfig", iface.Name, "hw", "ether", newMAC).Run()
	if err != nil {
		return err
	}
	return nil
}

func enableInterface(iface *net.Interface) error {
	err := exec.Command("ifconfig", iface.Name, "up").Run()
	if err != nil {
		return err
	}
	return nil
}
