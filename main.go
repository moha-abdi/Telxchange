package main

import (
	"fmt"

	"github.com/moha-abdi/telxchange/api"
	"github.com/moha-abdi/telxchange/config"
)

func main() {
	apiCli := api.NewApiClient(config.MFS_PROXY_25263)

	loginResponse, err := apiCli.Login(
		"3737373773",
		"7888",
		"hdjdjd",
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("Server Response: ", loginResponse.ServiceInfo.ResponseAttributes.ReplyMessage)

	if loginResponse.ServiceInfo.ResponseAttributes.ResultCode == "6020" {
		codeResponse, err := apiCli.Request2FA(
			loginResponse.ServiceInfo.ResponseAttributes.Username,
			"djddj",
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
			"djjdjjd",
			otpCode,
		)
		if err != nil {
			panic(err)
		}
		fmt.Println(result)

	}
}
