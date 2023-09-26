# Google Chrome Saved Passwords Extractor

This Go program allows you to extract saved passwords from Google Chrome's password database file. It can be used on both Windows and Linux operating systems.

## Prerequisites

Before you can use this program, make sure you have the following prerequisites:

- Go programming language installed on your system.
- The `github.com/mattn/go-sqlite3` package installed for Go. You can install it using the following command:
  ```bash
  go get github.com/mattn/go-sqlite3
    ```
## Usage

1. Clone this repository or download the main.go file to your local machine.
2. Open a terminal and navigate to the directory where the main.go file is located.
3. Run the program using the go run command:
```bash
go run main.go
```
4. The program will copy the Google Chrome password database file to the current directory with the name passwords_copy.db. It will then connect to this database and extract the saved passwords.

5. The extracted saved passwords will be displayed in the terminal.

## Attention!

The unethical or illegal use of this program is strictly prohibited. This code is for educational purposes only and should only be used for lawful purposes.
