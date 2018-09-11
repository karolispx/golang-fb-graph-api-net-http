package main

import (
	"fmt"

	fb "github.com/huandu/facebook"
)

const (
	FB_ACCESS_TOKEN = ""
	FB_API_VERSION  = "v3.1"
)

func main() {
	// Set up FB graphAPI version
	fb.Version = FB_API_VERSION

	// Run the sync
	printMessage("Getting accounts...")
	getAccounts()
}

// Get accounts
func getAccounts() {
	// Make request to /me/adaccounts to get accounts
	// Set your own limit or additioal parameters
	res, _ := fb.Get("/me/adaccounts?limit=2", fb.Params{
		"fields":       "id,name",
		"access_token": FB_ACCESS_TOKEN,
	})

	var items []fb.Result

	err := res.DecodeField("data", &items)

	if err != nil {
		// Error has happened when retrieving the data - repeat the call
		printMessage("An error has occurred while getting accounts. Repeating the call...")
		fmt.Printf("An error has happened %v", err)

		getAccounts()
	} else {
		// Accounts have been retrieved - save them to DB
		printMessage("Accounts have been retrieved successfully.")

		// For each item
		for _, item := range items {
			accountID := item["id"].(string)
			accountName := item["name"].(string)

			fmt.Println("Account ID: " + accountID + ", account name: " + accountName)
		}
	}
}

// Helper function for handling messages
func printMessage(message string) {
	fmt.Println("")
	fmt.Println(message)
	fmt.Println("")
}

// Helper function for handling errors
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
