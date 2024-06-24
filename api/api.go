package api

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/moha-abdi/telxchange/api/types/requests"
	"github.com/moha-abdi/telxchange/api/types/responses"
)

type APIClient struct {
	BaseURL    string
	httpClient *http.Client
}

func NewApiClient(baseURL string) *APIClient {
	f, err := os.OpenFile("C:/Users/lenovo/golog.log", os.O_APPEND|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	return &APIClient{
		BaseURL: baseURL,
		httpClient: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					KeyLogWriter: f,
				},
			},
		},
	}
}

func (c *APIClient) CreateToken(username string) (token string, err error) {
	url := c.BaseURL + "/api/createToken"
	client := c.httpClient

	tRequest := requests.NewTokenRequest()
	tRequest.ServiceInfo.RequestAttributes.Username = username

	alteredData, err := json.Marshal(&tRequest)
	if err != nil {
		return "", err
	}

	response, err := client.Post(url, "application/json", bytes.NewBuffer(alteredData))
	if err != nil {
		return "", err
	}

	resBody, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	var tResponse responses.TokenResponse
	err = json.Unmarshal(resBody, &tResponse)
	if err != nil {
		return "", err
	}

	return tResponse.ServiceInfo.ResponseAttributes.Token, nil
}

func (c *APIClient) Login(
	username string,
	password string,
	deviceId string,
) (response responses.LoginResponse, err error) {
	url := c.BaseURL + "/api/login"
	client := c.httpClient

	loginRequest := requests.NewLoginRequest()
	loginRequest.ServiceInfo.RequestAttributes.Username = username
	loginRequest.ServiceInfo.RequestAttributes.UserPassword = password
	loginRequest.ServiceInfo.RequestAttributes.DeviceId = deviceId

	data, err := json.Marshal(&loginRequest)
	if err != nil {
		return responses.LoginResponse{}, err
	}

	requestResponse, err := client.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return responses.LoginResponse{}, err
	}

	responseBody, err := io.ReadAll(requestResponse.Body)
	if err != nil {
		return responses.LoginResponse{}, err
	}

	var lResponse responses.LoginResponse
	if err = json.Unmarshal(responseBody, &lResponse); err != nil {
		return responses.LoginResponse{}, err
	}

	return lResponse, nil
}

func (c *APIClient) Request2FA(
	username string,
	deviceId string,
) (responses.VerificationCodeResponse, error) {
	url := c.BaseURL + "/api/requestVerificationCode"
	client := c.httpClient

	request := requests.NewVerificationCodeRequest()
	request.ServiceInfo.RequestAttributes.MobileNumber = username
	request.ServiceInfo.RequestAttributes.DeviceId = deviceId

	data, err := json.Marshal(&request)
	if err != nil {
		return responses.VerificationCodeResponse{}, err
	}

	requestRespone, err := client.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return responses.VerificationCodeResponse{}, err
	}

	requestResponeBody, err := io.ReadAll(requestRespone.Body)
	if err != nil {
		return responses.VerificationCodeResponse{}, err
	}

	var response responses.VerificationCodeResponse
	if err := json.Unmarshal(requestResponeBody, &response); err != nil {
		return responses.VerificationCodeResponse{}, err
	}

	return response, nil
}

func (c *APIClient) AuthenticateDevice(
	username string,
	deviceId string,
	otpCode string,
) (responses.AuthDeviceResponse, error) {
	url := c.BaseURL + "/api/authenticateDevice"
	client := c.httpClient

	request := requests.NewAuthDeviceRequest()
	request.ServiceInfo.RequestAttributes.MobileNumber = username
	request.ServiceInfo.RequestAttributes.DeviceId = deviceId
	request.ServiceInfo.RequestAttributes.OtpCode = otpCode

	data, err := json.Marshal(&request)
	if err != nil {
		return responses.AuthDeviceResponse{}, err
	}

	requestResponse, err := client.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return responses.AuthDeviceResponse{}, err
	}

	requestResponseBody, err := io.ReadAll(requestResponse.Body)
	if err != nil {
		return responses.AuthDeviceResponse{}, err
	}

	var response responses.AuthDeviceResponse
	if err := json.Unmarshal(requestResponseBody, &response); err != nil {
		return responses.AuthDeviceResponse{}, err
	}

	return response, nil
}
