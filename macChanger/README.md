# MAC Address Changer

## Overview

The MAC Address Changer is a simple command-line tool written in Go that allows you to change the MAC (Media Access Control) address of a network interface on a Linux system. This can be useful for privacy, security, or network troubleshooting purposes.

## Prerequisites

Before using this tool, ensure you have the following:

- **Go Programming Language**: You need Go installed on your system to compile and run this tool.

- **Administrative Privileges**: You typically need administrative privileges (e.g., root or sudo) to change the MAC address.

## Usage

To use the MAC Address Changer, follow these steps:

1. Clone or download the repository to your local machine.

2. Open a terminal and navigate to the directory where the `macChanger.go` file is located.

3. Run the program with the following command:

    ```bash
    go run macChanger.go -i [Network Interface Name] -m [New MAC Address]
    ```

   Replace `[Network Interface Name]` with the name of your network interface (e.g., eth0) and `[New MAC Address]` with the MAC address you want to set.

   Example:

    ```bash
    go run macChanger.go -i eth0 -m 00:11:22:33:44:55
    ```

4. The program will disable the network interface, set the new MAC address, and then re-enable the interface.

5. You will see a confirmation message indicating that the MAC address has been successfully changed.

## Attention!

The unethical or illegal use of this program is strictly prohibited. This code is for educational purposes only and should only be used for lawful purposes.
