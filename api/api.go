package api

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
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

type (
	ChannelName string
	AuthMode    string
)

const (
	Web       ChannelName = "Web"
	MobileApp ChannelName = "MobileApp"
)

const (
	Pin      AuthMode = "pin"
	Password AuthMode = "password"
)

const (
	DefaultChannel  = MobileApp
	DefaultAuthMode = Password
)

type BaseRequest struct {
	SchemaVersion string `json:"schemaVersion"`
	RequestID     string `json:"requestId"`
	Timestamp     int64  `json:"timestamp"`
	Channel       string `json:"channel"`
	SystemInfo    struct {
		SystemID     string `json:"systemId"`
		SystemSecret string `json:"systemSecret"`
	} `json:"systemInfo"`
}

type LocationInformation struct {
	CellID    string `json:"cellId"`
	LACId     string `json:"LACId"`
	MNC       string `json:"MNC"`
	MCC       string `json:"MCC"`
	IP        string `json:"IP"`
	MAC       string `json:"MAC"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

type TokenRequest struct {
	BaseRequest
	ServiceInfo struct {
		RequestAttributes struct {
			DeviceId             string              `json:"deviceId"`
			DeviceIdType         string              `json:"deviceIdType"`
			LocationInfo         LocationInformation `json:"locationInformation"`
			ReceiverLocalionInfo LocationInformation `json:"recieverlocationInformation"`
			Username             string              `json:"username"`
		} `json:"requestAttributes"`
		ServiceCode string `json:"serviceCode"`
		ServiceName string `json:"serviceName"`
	} `json:"serviceInfo"`
}

type TokenResponse struct {
	BaseRequest
	ServiceInfo struct {
		ResponseAttributes struct {
			Token                   string `json:"token"`
			EnablePayloadEncryption bool   `json:"enablePayloadEncryption"`
			ResultCode              int    `json:"resultCode"`
			ReplyMessage            string `json:"replyMessage"`
		} `json:"responseAttributes"`
	} `json:"serviceInfo"`
}

func (t *TokenRequest) setDefaults() {
	t.BaseRequest.SchemaVersion = "1.0"
	t.BaseRequest.RequestID = uuid.NewString()
	t.BaseRequest.Timestamp = time.Now().Unix()
	t.BaseRequest.Channel = string(DefaultChannel)
	t.ServiceInfo.ServiceCode = "0101"
	t.ServiceInfo.ServiceName = "CustomerLogin"
}

type LoginToken struct {
	BaseRequest
	ServiceInfo struct {
		RequestAttributes struct {
			DeviceId             string              `json:"deviceId"`
			DeviceIdType         string              `json:"deviceIdType"`
			LocationInfo         LocationInformation `json:"locationInformation"`
			ReceiverLocalionInfo LocationInformation `json:"recieverlocationInformation"`
			Username             string              `json:"username"`
			UserPassword         string              `json:"userPassword"`
			UserType             string              `json:"userType"`
			ChannelName          string              `json:"channelName"`
			ServiceType          string              `json:"serviceType"`
			AppVersion           string              `json:"appVersion,omitempty"`
			DeviceOS             string              `json:"deviceOS,omitempty"`
			AuthMode             string              `json:"authMode,omitempty"`
		} `json:"requestAttributes"`
		ServiceCode string `json:"serviceCode"`
		ServiceName string `json:"serviceName"`
	} `json:"serviceInfo"`
}

type LoginRequest struct {
	RequestId string      `json:"requestId"`
	Token     string      `json:"token"`
	Ltoken    *LoginToken `json:"-"`
}

type LoginResponse struct {
	BaseRequest
	ServiceInfo struct {
		ResponseAttributes struct {
			Username     string `json:"username"`
			UserType     string `json:"userType"`
			ResultCode   string `json:"resultCode"`
			ReplyMessage string `json:"replyMessage"`
		} `json:"responseAttributes"`
	} `json:"serviceInfo"`
}

func (l *LoginRequest) setDefaults() {
	if l.Ltoken == nil {
		l.Ltoken = &LoginToken{}
	}
	lt := l.Ltoken
	lt.BaseRequest.SchemaVersion = "1.0"
	lt.BaseRequest.RequestID = uuid.NewString()
	lt.BaseRequest.Timestamp = time.Now().Unix()
	lt.BaseRequest.Channel = string(DefaultChannel)
	lt.ServiceInfo.ServiceCode = "0101"
	lt.ServiceInfo.ServiceName = "CustomerLogin"
	lt.ServiceInfo.RequestAttributes.UserType = "CUSTOMER"
	lt.ServiceInfo.RequestAttributes.ChannelName = string(DefaultChannel)
	if DefaultChannel == MobileApp {
		lt.ServiceInfo.RequestAttributes.AppVersion = "8.2.2"
		lt.ServiceInfo.RequestAttributes.DeviceOS = "iOS"
		lt.ServiceInfo.RequestAttributes.AuthMode = string(Pin)
	}
	// Now set the Token and the RequestId in LoginRequest
	l.RequestId = lt.BaseRequest.RequestID
	res, err := json.Marshal(&lt)
	if err != nil {
		panic(err)
	}

	l.Token = string(res)
}

func (l *LoginRequest) setUsername(username string) {
	if l.Ltoken == nil {
		l.Ltoken = &LoginToken{}
	}
	l.Ltoken.ServiceInfo.RequestAttributes.Username = username
}

func (l *LoginRequest) setPassword(password string) {
	if l.Ltoken == nil {
		l.Ltoken = &LoginToken{}
	}
	l.Ltoken.ServiceInfo.RequestAttributes.UserPassword = password
}

func (c *APIClient) CreateToken(username string) (token string, err error) {
	url := c.BaseURL + "/api/createToken"
	client := c.httpClient

	var tRequest TokenRequest
	tRequest.setDefaults()
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

	var tResponse TokenResponse
	err = json.Unmarshal(resBody, &tResponse)
	if err != nil {
		return "", err
	}

	return tResponse.ServiceInfo.ResponseAttributes.Token, nil
}

func (c *APIClient) Login(username string, password string) (response LoginResponse, err error) {
	url := c.BaseURL + "/api/login"
	client := c.httpClient

	var loginRequest LoginRequest
	loginRequest.setUsername(username)
	loginRequest.setPassword(password)
	loginRequest.setDefaults()

	data, err := json.Marshal(&loginRequest)
	if err != nil {
		return LoginResponse{}, err
	}
	os.WriteFile("new.json", data, 0600)

	requestResponse, err := client.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return LoginResponse{}, err
	}

	responseBody, err := io.ReadAll(requestResponse.Body)
	if err != nil {
		return LoginResponse{}, err
	}

	fmt.Printf("Result is: %s", responseBody)

	var lResponse LoginResponse
	if err = json.Unmarshal(responseBody, &lResponse); err != nil {
		return LoginResponse{}, err
	}

	return lResponse, nil
}
