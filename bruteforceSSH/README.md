# SSH Password Cracker

This is a simple Go program that attempts to crack an SSH login by brute-forcing passwords using a provided wordlist. It connects to a target SSH server and iterates through the wordlist, trying each password until it either successfully logs in or exhausts the wordlist.

## Prerequisites

Before you begin, ensure you have met the following requirements:

- Go (Golang) is installed on your system.
- You have SSH access to the target server.
- You have a wordlist file containing passwords.

## Usage

1. Clone the repository or download the `main.go` file.

2. Build the Go program:

   ```bash
   go build main.go
   ./main -i <target_ip> -u <username> -w <wordlist_file>
   ```
   Optional flags:
    ```
    -p (Target Port Number, Default: 22)
    -t (Request timeout, Default: 5)
    ```

The program will start attempting to crack the SSH password. If successful, it will display the password and exit. If unsuccessful, it will continue until it either finds the password or exhausts the wordlist.

## Attention!

The unethical or illegal use of this program is strictly prohibited. This code is for educational purposes only and should only be used for lawful purposes.
