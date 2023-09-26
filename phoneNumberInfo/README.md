# Phone Number Validation

This Go program validates phone numbers using the [apilayer Number Verification API](https://apilayer.com/number-verification). The API allows you to verify the accuracy of a given phone number and retrieve relevant information.

## Prerequisites

Before running this program, ensure you have the following:

- Go installed on your system.
- An API key from [apilayer](https://apilayer.com/number-verification) to access the API.

## Installation

1. Clone this repository to your local machine:

```bash
Clone this repository or download the Go source code file.
cd phone-number-validation
go build
./phone-number-validation
```
1. You will be prompted to enter a phone number (e.g., 1234567890).

2. The program will use the apilayer Number Verification API to check the validity of the entered phone number.

3. The result of the API request will be displayed in the console.

## Configuration

Before running the program, open the main.go file and replace "your-api-key" with your actual apilayer API key.

```bash
req.Header.Set("apikey", "your-api-key") // Replace with your API key
```

## Attention!

The unethical or illegal use of this program is strictly prohibited. This code is for educational purposes only and should only be used for lawful purposes.