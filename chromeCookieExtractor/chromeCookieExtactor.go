package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
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

	var cookiePath string
	var isWindows bool

	if osName := os.Getenv("OS"); osName == "Windows_NT" {
		cookiePath = filepath.Join(username, "AppData", "Local", "Google", "Chrome", "User Data", "Default", "Network", "Cookies")
		isWindows = true
	} else {
		cookiePath = filepath.Join("/home", username, ".config", "google-chrome", "Default", "Cookies")
		isWindows = false
	}

	destDBPath := "cookies_copy.db"

	sourceData, err := ioutil.ReadFile(cookiePath)
	if err != nil {
		log.Fatalf("Failed to read the source database file: %v", err)
	}

	err = ioutil.WriteFile(destDBPath, sourceData, 0644)
	if err != nil {
		log.Fatalf("Failed to write to the destination database file: %v", err)
	}

	fmt.Printf("Database copied successfully: %s\n", destDBPath)

	db, err := sql.Open("sqlite3", destDBPath)
	if err != nil {
		log.Fatalf("Failed to connect to the cookie database: %v", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT name, value, host_key FROM cookies")
	if err != nil {
		log.Fatalf("Failed to read cookies: %v", err)
	}
	defer rows.Close()

	fmt.Println("Chrome Cookies:")
	for rows.Next() {
		var name, value, hostKey string
		if err := rows.Scan(&name, &value, &hostKey); err != nil {
			log.Fatalf("Failed to read cookie data: %v", err)
		}
		fmt.Printf("Cookie Name: %s, Value: %s, Host Key: %s\n", name, value, hostKey)
	}

	if isWindows {
		fmt.Println("This is a Windows operating system.")
	} else {
		fmt.Println("This is a Linux operating system.")
	}
}
