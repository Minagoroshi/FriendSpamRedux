//go:generate goversioninfo -icon=ico.ico

package main

import (
	"FriendSpamRedux/cmd"
	"FriendSpamRedux/requests"
	"FriendSpamRedux/utils"
	"fmt"
	"github.com/Delta456/box-cli-maker/v2"
	"github.com/gookit/color"
	"log"
	"os"
	"strconv"
)

func main() {
	//Load Proxies
	proxyList := cmd.LoadProxies()
	if len(proxyList) == 0 {
		color.Red.Println("No proxies found. Please add some proxies to your proxies file.")
	}
	proxyNum := strconv.Itoa(len(proxyList))

	//Load Accounts
	accounts := cmd.LoadAccounts()
	if len(accounts) == 0 {
		color.Red.Println("No accounts found. Please add some accounts to your accounts file.")
	}
	accNum := strconv.Itoa(len(accounts))

	color.Cyan.Println("Welcome to FriendSpamRedux!")

	for {
		//Use the box-cli-maker to create a box-cli
		Box := box.New(box.Config{Px: 2, Py: 1, Type: "Single", Color: "HiGreen"})
		Box.Print("VRChat Friend Spam Redux\n"+
			"by Top",
			"1. Friend Spam (UserID)\n"+
				//	"2. Set Avatar\n"+
				"3. Exit\n\n"+
				"Accounts Loaded: "+accNum+"\n"+
				"Proxies Loaded: "+proxyNum)

		//Get the user input
		var input string
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println(err)
		}

		//Check the input
		switch input {
		case "1": //Friend Request
			utils.ClearScreen()

			//Prompt for proxies
			fmt.Println("Use Proxies? (Warning! Without proxies you will get rate-limited over 4 accounts, and risk bans!) (y/n)")
			var promptProxy string
			_, err := fmt.Scanln(&promptProxy)
			if err != nil {
				fmt.Println(err)
			}

			//Take user input for userID
			fmt.Print("Input UserID: ")
			var userID string
			_, err = fmt.Scanln(&userID)
			if err != nil {
				fmt.Println(err)
			}

			//Begin Friend Spam
			log.Println("Sending Friend Requests")

			for i, account := range accounts {
				//Check if the account is valid
				if !cmd.CheckAccounts(account) {
					color.Error.Println("Account: " + strconv.Itoa(i) + " is invalid. Please check your accounts file.")
					continue
				}

				//Load proxies and login
				var proxy string
				var useProxy bool
				if promptProxy == "y" {
					proxy = cmd.GetProxy()
					valid := cmd.CheckProxy(proxy)
					if !valid {
						color.Error.Println("Proxy: " + proxy + " is invalid. Please check your proxies file.")
						continue
					}
					useProxy = true
				} else {
					useProxy = false
				}

				//Login to VRChat and get the cookies
				loginResp, err, loginErr := requests.Login(account, proxy, useProxy)
				if loginErr != nil {
					color.Error.Println("Account: " + strconv.Itoa(i) + " failed to login. Please check your accounts file.")
					continue
				}
				if err != nil {
					color.Error.Println("Error: " + err.Error())
					continue
				}

				cookies := loginResp.RawResponse.Cookies()

				//Send the friend request
				resp, err := requests.FriendRequest(userID, cookies, proxy, useProxy)
				if err != nil {
					fmt.Println(err)
					continue
				}

				//Check the response
				if resp.Result().(*requests.FriendRequestResponse).Message == "" || resp.RawResponse.StatusCode == 200 {
					color.HiGreen.Println("Friend Request Sent!", i+1, "of", len(accounts))
				} else {
					color.HiRed.Println("Friend Request Failed!")
				}

			}
		//case "2": //Set Avatar
		//	utils.ClearScreen()
		//	log.Println("Setting Avatar")
		case "3": //Exit
			utils.ClearScreen()
			log.Println("Exiting")
			os.Exit(0)
		default:
			utils.ClearScreen()
			log.Println("Invalid Input")
		}

	}

}
