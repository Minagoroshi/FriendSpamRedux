package cmd

import (
	"bufio"
	"log"
	"os"
	"regexp"
)

var accountRegex = "^[a-zA-Z0-9]{3,20}:[a-zA-Z0-9]{3,20}$"

//LoadAccounts is a function to load each line of 'appconfig/accounts' into a slice of strings
func LoadAccounts() []string {
	var accounts []string
	file, err := os.Open("appconfig/accounts")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		accounts = append(accounts, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return accounts
}

//CheckAccounts is a function to check if the account is valid using the accountRegex regex
func CheckAccounts(account string) bool {
	valid, _ := regexp.MatchString(accountRegex, account)
	return valid
}
