package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	var phoneNumber string

	fmt.Print("Enter the phone number (e.g., 1234567890): ")
	fmt.Scan(&phoneNumber)

	url := fmt.Sprintf("https://api.apilayer.com/number_verification/validate?number=%s", phoneNumber)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Set("apikey", "your-api-key") // Replace with your API key

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("API response failed:", res.Status)
		return
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("API response could not be read:", err)
		return
	}

	fmt.Println(string(body))
}
