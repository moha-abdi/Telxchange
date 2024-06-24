package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/moha-abdi/telxchange/api"
	"github.com/moha-abdi/telxchange/config"
)

func main() {
	// Load .env file before everything else
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	username := os.Getenv("TELXCHANGE_USERNAME")
	password := os.Getenv("TELXCHANGE_PASSWORD")
	device_id := os.Getenv("TELXCHANGE_DEVICE_ID")
	fmt.Println("Loaded account Info ->", "Username:", username, "Device ID:", device_id)

	apiCli := api.NewApiClient(config.MFS_PROXY_25263)

	loginResponse, err := apiCli.Login(
		username,
		password,
		device_id,
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("Server Response: ", loginResponse.ServiceInfo.ResponseAttributes.ReplyMessage)

	if loginResponse.ServiceInfo.ResponseAttributes.ResultCode == "6020" {
		codeResponse, err := apiCli.Request2FA(
			loginResponse.ServiceInfo.ResponseAttributes.Username,
			device_id,
		)
		if err != nil {
			panic(err)
		}
		fmt.Println(codeResponse)

		fmt.Print("Please enter the otp Code: ")
		var otpCode string
		fmt.Scan(&otpCode)

		result, err := apiCli.AuthenticateDevice(
			loginResponse.ServiceInfo.ResponseAttributes.Username,
			device_id,
			otpCode,
		)
		if err != nil {
			panic(err)
		}
		fmt.Println(result)

	}
}
