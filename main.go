package main

import (
	"fmt"
	"moha/telxchange/api"
)

func main() {
	apiCli := api.NewApiClient("https://myaccount.telesom.com")
	token, err := apiCli.CreateToken("25848484")
	fmt.Print(token)
	if err != nil {
		panic(err)
	}
	loginResponse, err := apiCli.Login("252383838", "8484")
	if err != nil {
		panic(err)
	}
	fmt.Println("Server Response: ", loginResponse.ServiceInfo.ResponseAttributes.ReplyMessage)
}
